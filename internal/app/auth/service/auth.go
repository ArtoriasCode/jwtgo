package service

import (
	"context"
	service2 "jwtgo/internal/pkg/interface/service"

	"jwtgo/internal/app/auth/server/grpc/dto"
	"jwtgo/internal/app/auth/server/grpc/mapper"
	customErr "jwtgo/internal/pkg/error"
	userPb "jwtgo/internal/pkg/proto/user"
	"jwtgo/pkg/logging"
)

type AuthService struct {
	userService     userPb.UserServiceClient
	jwtService      service2.JWTService
	passwordService service2.PasswordService
	logger          *logging.Logger
}

func NewAuthService(
	userService userPb.UserServiceClient,
	jwtService service2.JWTService,
	passwordService service2.PasswordService,
	logger *logging.Logger,
) *AuthService {
	return &AuthService{
		userService:     userService,
		jwtService:      jwtService,
		passwordService: passwordService,
		logger:          logger,
	}
}

func (s *AuthService) SignUp(ctx context.Context, userCredentialsDTO *dto.UserCredentialsDTO) (bool, error) {
	getByEmailRequest := mapper.MapEmailToUserGetByEmailRequest(userCredentialsDTO.Email)

	getByEmailResponse, err := s.userService.GetByEmail(ctx, getByEmailRequest)
	if err != nil {
		s.logger.Error("Error while getting user: ", err)
		return false, customErr.NewInternalServerError("Failed to check user email")
	}

	if getByEmailResponse.User != nil {
		return false, customErr.NewAlreadyExistsError("Email already exists")
	}

	localSalt, err := s.passwordService.GenerateSalt(32)
	if err != nil {
		s.logger.Error("Error while generating local salt: ", err)
		return false, customErr.NewInternalServerError("Failed to create a user")
	}

	hashedPassword, err := s.passwordService.HashPassword(userCredentialsDTO.Password, localSalt)
	if err != nil {
		s.logger.Error("Error while hashing password: ", err)
		return false, customErr.NewInternalServerError("Failed to create a user")
	}

	userCredentialsDTO.Password = hashedPassword

	createRequest := mapper.MapUserCredentialsDTOToUserCreateRequest(userCredentialsDTO)
	createRequest.Salt = localSalt

	_, err = s.userService.Create(ctx, createRequest)
	if err != nil {
		s.logger.Error("Error while creating user: ", err)
		return false, customErr.NewInternalServerError("Failed to create a user")
	}

	return true, nil
}

func (s *AuthService) SignIn(ctx context.Context, userCredentialsDTO *dto.UserCredentialsDTO) (*dto.UserTokensDTO, error) {
	getByEmailRequest := mapper.MapEmailToUserGetByEmailRequest(userCredentialsDTO.Email)

	getByEmailResponse, err := s.userService.GetByEmail(ctx, getByEmailRequest)
	if err != nil {
		s.logger.Error("Error while getting user: ", err)
		return nil, customErr.NewInternalServerError("Failed to check user email")
	}

	if getByEmailResponse.User == nil {
		return nil, customErr.NewInvalidCredentialsError("Invalid login or password")
	}

	passwordIsValid := s.passwordService.VerifyPassword(userCredentialsDTO.Password, getByEmailResponse.User.Password, getByEmailResponse.User.Salt)
	if !passwordIsValid {
		return nil, customErr.NewInvalidCredentialsError("Invalid login or password")
	}

	accessToken, refreshToken, err := s.jwtService.GenerateTokens(getByEmailResponse.User.Id)
	if err != nil {
		s.logger.Error("Error while generating tokens: ", err)
		return nil, customErr.NewInternalServerError("Token generation error")
	}

	getByEmailResponse.User.RefreshToken = refreshToken
	updateRequest := mapper.MapUserGetByEmailResponseToUserUpdateRequest(getByEmailResponse)

	_, err = s.userService.Update(ctx, updateRequest)
	if err != nil {
		s.logger.Error("Error while updating user: ", err)
		return nil, customErr.NewInternalServerError("Token updating error")
	}

	return mapper.MapTokensToUserTokensDTO(accessToken, refreshToken), nil
}

func (s *AuthService) SignOut(ctx context.Context, refreshTokenDTO *dto.UserTokenDTO) (bool, error) {
	claims, err := s.jwtService.ValidateToken(refreshTokenDTO.Token)
	if err != nil {
		return false, err
	}

	getByIdRequest := mapper.MapIdToUserGetByIdRequest(claims.Id)

	getByIdResponse, err := s.userService.GetById(ctx, getByIdRequest)
	if err != nil {
		s.logger.Error("Error while getting user: ", err)
		return false, customErr.NewInternalServerError("Failed to check user id")
	}

	if getByIdResponse == nil {
		return false, customErr.NewNotFoundError("User not found")
	}

	getByIdResponse.User.RefreshToken = ""
	getByIdResponse.User.Id = claims.Id

	updateRequest := mapper.MapUserGetByIdResponseToUserUpdateRequest(getByIdResponse)

	_, err = s.userService.Update(ctx, updateRequest)
	if err != nil {
		s.logger.Error("Error while updating user: ", err)
		return false, customErr.NewInternalServerError("User updating error")
	}

	return true, nil
}

func (s *AuthService) Refresh(ctx context.Context, refreshTokenDTO *dto.UserTokenDTO) (*dto.UserTokensDTO, error) {
	claims, err := s.jwtService.ValidateToken(refreshTokenDTO.Token)
	if err != nil {
		return nil, err
	}

	getByIdRequest := mapper.MapIdToUserGetByIdRequest(claims.Id)

	getByIdResponse, err := s.userService.GetById(ctx, getByIdRequest)
	if err != nil {
		s.logger.Error("Error while getting user: ", err)
		return nil, customErr.NewInternalServerError("Failed to check user id")
	}

	if getByIdResponse == nil {
		return nil, customErr.NewNotFoundError("User not found")
	}

	if refreshTokenDTO.Token != getByIdResponse.User.RefreshToken {
		return nil, customErr.NewInvalidTokenError("Invalid refresh token")
	}

	accessToken, refreshToken, err := s.jwtService.GenerateTokens(getByIdResponse.User.Id)
	if err != nil {
		s.logger.Error("Error while generating tokens: ", err)
		return nil, customErr.NewInternalServerError("Token generation error")
	}

	getByIdResponse.User.RefreshToken = refreshToken
	getByIdResponse.User.Id = claims.Id

	updateRequest := mapper.MapUserGetByIdResponseToUserUpdateRequest(getByIdResponse)

	_, err = s.userService.Update(ctx, updateRequest)
	if err != nil {
		s.logger.Error("Error while updating user: ", err)
		return nil, customErr.NewInternalServerError("Token updating error")
	}

	return mapper.MapTokensToUserTokensDTO(accessToken, refreshToken), nil
}
