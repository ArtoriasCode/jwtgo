package service

import (
	"context"

	"jwtgo/internal/app/auth/controller/grpc/dto"
	"jwtgo/internal/app/auth/controller/grpc/mapper"
	authServiceIface "jwtgo/internal/app/auth/interface/service"
	customErr "jwtgo/internal/pkg/error/type"
	pkgServiceIface "jwtgo/internal/pkg/interface/service"
	userPb "jwtgo/internal/pkg/proto/user"
	"jwtgo/pkg/logging"
)

type AuthService struct {
	userMicroService userPb.UserServiceClient
	jwtService       pkgServiceIface.JWTServiceIface
	passwordService  authServiceIface.PasswordServiceIface
	logger           *logging.Logger
}

func NewAuthService(
	userMicroService userPb.UserServiceClient,
	jwtService pkgServiceIface.JWTServiceIface,
	passwordService authServiceIface.PasswordServiceIface,
	logger *logging.Logger,
) *AuthService {
	return &AuthService{
		userMicroService: userMicroService,
		jwtService:       jwtService,
		passwordService:  passwordService,
		logger:           logger,
	}
}

func (s *AuthService) SignUp(ctx context.Context, userCredentialsDTO *dto.UserCredentialsDTO) (bool, customErr.BaseErrorIface) {
	getByEmailRequest := mapper.MapEmailToUserGetByEmailRequest(userCredentialsDTO.Email)

	getByEmailResponse, err := s.userMicroService.GetByEmail(ctx, getByEmailRequest)
	if err != nil {
		s.logger.Error("[AuthService -> SignUp -> GetByEmail]: ", err)
		return false, customErr.NewInternalServerError("Failed to create user")
	}

	if getByEmailResponse.User != nil {
		return false, customErr.NewAlreadyExistsError("Email already exists")
	}

	localSalt, err := s.passwordService.GenerateSalt(32)
	if err != nil {
		s.logger.Error("[AuthService -> SignUp -> GenerateSalt]: ", err)
		return false, customErr.NewInternalServerError("Failed to create user")
	}

	hashedPassword, err := s.passwordService.HashPassword(userCredentialsDTO.Password, localSalt)
	if err != nil {
		s.logger.Error("[AuthService -> SignUp -> HashPassword]: ", err)
		return false, customErr.NewInternalServerError("Failed to create user")
	}

	userCredentialsDTO.Password = hashedPassword

	createRequest := mapper.MapUserCredentialsDTOToUserCreateRequest(userCredentialsDTO)
	createRequest.Security.Salt = localSalt

	_, err = s.userMicroService.Create(ctx, createRequest)
	if err != nil {
		s.logger.Error("[AuthService -> SignUp -> Create]: ", err)
		return false, customErr.NewInternalServerError("Failed to create user")
	}

	return true, nil
}

func (s *AuthService) SignIn(ctx context.Context, userCredentialsDTO *dto.UserCredentialsDTO) (*dto.UserTokensDTO, customErr.BaseErrorIface) {
	getByEmailRequest := mapper.MapEmailToUserGetByEmailRequest(userCredentialsDTO.Email)

	getByEmailResponse, err := s.userMicroService.GetByEmail(ctx, getByEmailRequest)
	if err != nil {
		s.logger.Error("[AuthService -> SignIn -> GetByEmail]: ", err)
		return nil, customErr.NewInternalServerError("Failed to sign in user")
	}

	if getByEmailResponse.User == nil {
		return nil, customErr.NewInvalidCredentialsError("Invalid login or password")
	}

	passwordIsValid := s.passwordService.VerifyPassword(
		userCredentialsDTO.Password,
		getByEmailResponse.User.Security.Password,
		getByEmailResponse.User.Security.Salt,
	)

	if !passwordIsValid {
		return nil, customErr.NewInvalidCredentialsError("Invalid login or password")
	}

	accessToken, refreshToken, err := s.jwtService.GenerateTokens(getByEmailResponse.User.Id, getByEmailResponse.User.Role)
	if err != nil {
		s.logger.Error("[AuthService -> SignIn -> GenerateTokens]: ", err)
		return nil, customErr.NewInternalServerError("Failed to sign in user")
	}

	getByEmailResponse.User.Security.RefreshToken = refreshToken
	updateRequest := mapper.MapUserGetByEmailResponseToUserUpdateRequest(getByEmailResponse)

	_, err = s.userMicroService.Update(ctx, updateRequest)
	if err != nil {
		s.logger.Error("[AuthService -> SignIn -> Update]: ", err)
		return nil, customErr.NewInternalServerError("Failed to sign in user")
	}

	return mapper.MapTokensToUserTokensDTO(accessToken, refreshToken), nil
}

func (s *AuthService) SignOut(ctx context.Context, refreshTokenDTO *dto.UserTokenDTO) (bool, customErr.BaseErrorIface) {
	claims, err := s.jwtService.ValidateToken(refreshTokenDTO.Token)
	if err != nil {
		return false, err
	}

	getByIdRequest := mapper.MapIdToUserGetByIdRequest(claims.Id)

	getByIdResponse, err := s.userMicroService.GetById(ctx, getByIdRequest)
	if err != nil {
		s.logger.Error("[AuthService -> SignOut -> GetById]: ", err)
		return false, customErr.NewInternalServerError("Failed to sign out user")
	}

	if getByIdResponse == nil {
		return false, customErr.NewNotFoundError("User not found")
	}

	getByIdResponse.User.Security.RefreshToken = ""
	getByIdResponse.User.Id = claims.Id

	updateRequest := mapper.MapUserGetByIdResponseToUserUpdateRequest(getByIdResponse)

	_, err = s.userMicroService.Update(ctx, updateRequest)
	if err != nil {
		s.logger.Error("[AuthService -> SignOut -> Update]: ", err)
		return false, customErr.NewInternalServerError("Failed to sign out user")
	}

	return true, nil
}

func (s *AuthService) Refresh(ctx context.Context, refreshTokenDTO *dto.UserTokenDTO) (*dto.UserTokensDTO, customErr.BaseErrorIface) {
	claims, err := s.jwtService.ValidateToken(refreshTokenDTO.Token)
	if err != nil {
		return nil, err
	}

	getByIdRequest := mapper.MapIdToUserGetByIdRequest(claims.Id)

	getByIdResponse, err := s.userMicroService.GetById(ctx, getByIdRequest)
	if err != nil {
		s.logger.Error("[AuthService -> Refresh -> GetById]: ", err)
		return nil, customErr.NewInternalServerError("Failed to refresh tokens")
	}

	if getByIdResponse == nil {
		return nil, customErr.NewNotFoundError("User not found")
	}

	if refreshTokenDTO.Token != getByIdResponse.User.Security.RefreshToken {
		return nil, customErr.NewInvalidTokenError("Invalid refresh token")
	}

	accessToken, refreshToken, err := s.jwtService.GenerateTokens(getByIdResponse.User.Id, getByIdResponse.User.Role)
	if err != nil {
		s.logger.Error("[AuthService -> Refresh -> GenerateTokens]: ", err)
		return nil, customErr.NewInternalServerError("Failed to refresh tokens")
	}

	getByIdResponse.User.Security.RefreshToken = refreshToken
	getByIdResponse.User.Id = claims.Id

	updateRequest := mapper.MapUserGetByIdResponseToUserUpdateRequest(getByIdResponse)

	_, err = s.userMicroService.Update(ctx, updateRequest)
	if err != nil {
		s.logger.Error("[AuthService -> Refresh -> Update]: ", err)
		return nil, customErr.NewInternalServerError("Failed to refresh tokens")
	}

	return mapper.MapTokensToUserTokensDTO(accessToken, refreshToken), nil
}
