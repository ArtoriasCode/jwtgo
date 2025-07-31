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

func (s *AuthService) SignUp(ctx context.Context, signUpRequestDTO *dto.SignUpRequestDTO) (bool, customErr.BaseErrorIface) {
	userGetByEmailRequest := mapper.MapEmailToUserGetByEmailRequest(signUpRequestDTO.Email)

	userGetByEmailResponse, err := s.userMicroService.GetByEmail(ctx, userGetByEmailRequest)
	if err != nil {
		s.logger.Error("[AuthService -> SignUp -> GetByEmail]: ", err)
		return false, customErr.NewInternalServerError("Failed to create user")
	}

	if userGetByEmailResponse.User != nil {
		return false, customErr.NewAlreadyExistsError("Email already exists")
	}

	localSalt, err := s.passwordService.GenerateSalt(32)
	if err != nil {
		s.logger.Error("[AuthService -> SignUp -> GenerateSalt]: ", err)
		return false, customErr.NewInternalServerError("Failed to create user")
	}

	hashedPassword, err := s.passwordService.HashPassword(signUpRequestDTO.Password, localSalt)
	if err != nil {
		s.logger.Error("[AuthService -> SignUp -> HashPassword]: ", err)
		return false, customErr.NewInternalServerError("Failed to create user")
	}

	signUpRequestDTO.Password = hashedPassword

	userCreateRequest := mapper.MapSignUpRequestDTOToUserCreateRequest(signUpRequestDTO)
	userCreateRequest.Security.Salt = localSalt

	_, err = s.userMicroService.Create(ctx, userCreateRequest)
	if err != nil {
		s.logger.Error("[AuthService -> SignUp -> Create]: ", err)
		return false, customErr.NewInternalServerError("Failed to create user")
	}

	return true, nil
}

func (s *AuthService) SignIn(ctx context.Context, signInRequestDTO *dto.SignInRequestDTO) (*dto.UserTokensDTO, customErr.BaseErrorIface) {
	userGetByEmailRequest := mapper.MapEmailToUserGetByEmailRequest(signInRequestDTO.Email)

	userGetByEmailResponse, err := s.userMicroService.GetByEmail(ctx, userGetByEmailRequest)
	if err != nil {
		s.logger.Error("[AuthService -> SignIn -> GetByEmail]: ", err)
		return nil, customErr.NewInternalServerError("Failed to sign in user")
	}

	if userGetByEmailResponse.User == nil {
		return nil, customErr.NewInvalidCredentialsError("Invalid login or password")
	}

	passwordIsValid := s.passwordService.VerifyPassword(
		signInRequestDTO.Password,
		userGetByEmailResponse.User.Security.Password,
		userGetByEmailResponse.User.Security.Salt,
	)

	if !passwordIsValid {
		return nil, customErr.NewInvalidCredentialsError("Invalid login or password")
	}

	accessToken, refreshToken, err := s.jwtService.GenerateTokens(userGetByEmailResponse.User.Id, userGetByEmailResponse.User.Role)
	if err != nil {
		s.logger.Error("[AuthService -> SignIn -> GenerateTokens]: ", err)
		return nil, customErr.NewInternalServerError("Failed to sign in user")
	}

	userGetByEmailResponse.User.Security.RefreshToken = refreshToken
	updateRequest := mapper.MapUserGetByEmailResponseToUserUpdateRequest(userGetByEmailResponse)

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

	userGetByIdRequest := mapper.MapIdToUserGetByIdRequest(claims.Id)

	userGetByIdResponse, err := s.userMicroService.GetById(ctx, userGetByIdRequest)
	if err != nil {
		s.logger.Error("[AuthService -> SignOut -> GetById]: ", err)
		return false, customErr.NewInternalServerError("Failed to sign out user")
	}

	if userGetByIdResponse.User == nil {
		return false, customErr.NewNotFoundError("User not found")
	}

	userGetByIdResponse.User.Security.RefreshToken = ""
	userGetByIdResponse.User.Id = claims.Id

	userUpdateRequest := mapper.MapUserGetByIdResponseToUserUpdateRequest(userGetByIdResponse)

	_, err = s.userMicroService.Update(ctx, userUpdateRequest)
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

	userGetByIdRequest := mapper.MapIdToUserGetByIdRequest(claims.Id)

	userGetByIdResponse, err := s.userMicroService.GetById(ctx, userGetByIdRequest)
	if err != nil {
		s.logger.Error("[AuthService -> Refresh -> GetById]: ", err)
		return nil, customErr.NewInternalServerError("Failed to refresh tokens")
	}

	if userGetByIdResponse.User == nil {
		return nil, customErr.NewNotFoundError("User not found")
	}

	if refreshTokenDTO.Token != userGetByIdResponse.User.Security.RefreshToken {
		return nil, customErr.NewInvalidTokenError("Invalid refresh token")
	}

	accessToken, refreshToken, err := s.jwtService.GenerateTokens(userGetByIdResponse.User.Id, userGetByIdResponse.User.Role)
	if err != nil {
		s.logger.Error("[AuthService -> Refresh -> GenerateTokens]: ", err)
		return nil, customErr.NewInternalServerError("Failed to refresh tokens")
	}

	userGetByIdResponse.User.Security.RefreshToken = refreshToken
	userGetByIdResponse.User.Id = claims.Id

	userUpdateRequest := mapper.MapUserGetByIdResponseToUserUpdateRequest(userGetByIdResponse)

	_, err = s.userMicroService.Update(ctx, userUpdateRequest)
	if err != nil {
		s.logger.Error("[AuthService -> Refresh -> Update]: ", err)
		return nil, customErr.NewInternalServerError("Failed to refresh tokens")
	}

	return mapper.MapTokensToUserTokensDTO(accessToken, refreshToken), nil
}
