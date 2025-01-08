package service

import (
	"context"

	"jwtgo/internal/app/controller/http/dto"
	"jwtgo/internal/app/controller/http/mapper"
	customErr "jwtgo/internal/app/error"
	repositoryInterface "jwtgo/internal/app/interface/repository"
	serviceInterface "jwtgo/internal/app/interface/service"
	"jwtgo/pkg/logging"
)

type AuthService struct {
	userRepository  repositoryInterface.UserRepository
	jwtService      serviceInterface.JWTService
	passwordService serviceInterface.PasswordService
	logger          *logging.Logger
}

func NewAuthService(
	userRepository repositoryInterface.UserRepository,
	jwtService serviceInterface.JWTService,
	passwordService serviceInterface.PasswordService,
	logger *logging.Logger,
) *AuthService {
	return &AuthService{
		userRepository:  userRepository,
		jwtService:      jwtService,
		passwordService: passwordService,
		logger:          logger,
	}
}

func (s *AuthService) SignUp(ctx context.Context, userCredentialsDTO *dto.UserCredentialsDTO) (bool, error) {
	existingUserEntity, err := s.userRepository.GetByEmail(ctx, userCredentialsDTO.Email)
	if err != nil {
		s.logger.Error("Error while getting user: ", err)
		return false, customErr.NewInternalServerError("Failed to check user email")
	}

	if existingUserEntity != nil {
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

	userCreateEntity := mapper.MapUserCredentialsDTOToDomainUser(userCredentialsDTO)
	userCreateEntity.Salt = localSalt

	_, err = s.userRepository.Create(ctx, userCreateEntity)
	if err != nil {
		s.logger.Error("Error while creating user: ", err)
		return false, customErr.NewInternalServerError("Failed to create a user")
	}

	return true, nil
}

func (s *AuthService) SignIn(ctx context.Context, userCredentialsDTO *dto.UserCredentialsDTO) (*dto.UserTokensDTO, error) {
	existingUserEntity, err := s.userRepository.GetByEmail(ctx, userCredentialsDTO.Email)
	if err != nil {
		s.logger.Error("Error while getting user: ", err)
		return nil, customErr.NewInternalServerError("Failed to check user email")
	}

	if existingUserEntity == nil {
		return nil, customErr.NewInvalidCredentialsError("Invalid login or password")
	}

	passwordIsValid := s.passwordService.VerifyPassword(userCredentialsDTO.Password, existingUserEntity.Password, existingUserEntity.Salt)
	if !passwordIsValid {
		return nil, customErr.NewInvalidCredentialsError("Invalid login or password")
	}

	accessToken, refreshToken, err := s.jwtService.GenerateTokens(existingUserEntity.Id)
	if err != nil {
		s.logger.Error("Error while generating tokens: ", err)
		return nil, customErr.NewInternalServerError("Token generation error")
	}

	existingUserEntity.RefreshToken = refreshToken

	_, err = s.userRepository.Update(ctx, existingUserEntity.Id, existingUserEntity)
	if err != nil {
		s.logger.Error("Error while updating user: ", err)
		return nil, customErr.NewInternalServerError("Token updating error")
	}

	return mapper.MapToUserTokensDTO(accessToken, refreshToken), nil
}

func (s *AuthService) SignOut(ctx context.Context, refreshTokenDTO *dto.UserTokenDTO) error {
	claims, err := s.jwtService.ValidateToken(refreshTokenDTO.RefreshToken)
	if err != nil {
		return err
	}

	existingUserEntity, err := s.userRepository.GetById(ctx, claims.Id)
	if err != nil {
		s.logger.Error("Error while getting user: ", err)
		return customErr.NewInternalServerError("Failed to check user id")
	}

	if existingUserEntity == nil {
		return customErr.NewUserNotFoundError("User not found")
	}

	existingUserEntity.RefreshToken = ""

	_, err = s.userRepository.Update(ctx, claims.Id, existingUserEntity)
	if err != nil {
		s.logger.Error("Error while updating user: ", err)
		return customErr.NewInternalServerError("User updating error")
	}

	return nil
}

func (s *AuthService) Refresh(ctx context.Context, refreshTokenDTO *dto.UserTokenDTO) (*dto.UserTokensDTO, error) {
	claims, err := s.jwtService.ValidateToken(refreshTokenDTO.RefreshToken)
	if err != nil {
		return nil, err
	}

	existingUserEntity, err := s.userRepository.GetById(ctx, claims.Id)
	if err != nil {
		s.logger.Error("Error while getting user: ", err)
		return nil, customErr.NewInternalServerError("Failed to check user id")
	}

	if existingUserEntity == nil {
		return nil, customErr.NewUserNotFoundError("User not found")
	}

	if refreshTokenDTO.RefreshToken != existingUserEntity.RefreshToken {
		return nil, customErr.NewInvalidTokenError("Invalid refresh token")
	}

	accessToken, refreshToken, err := s.jwtService.GenerateTokens(existingUserEntity.Id)
	if err != nil {
		s.logger.Error("Error while generating tokens: ", err)
		return nil, customErr.NewInternalServerError("Token generation error")
	}

	existingUserEntity.RefreshToken = refreshToken

	_, err = s.userRepository.Update(ctx, claims.Id, existingUserEntity)
	if err != nil {
		s.logger.Error("Error while updating user: ", err)
		return nil, customErr.NewInternalServerError("Token updating error")
	}

	return mapper.MapToUserTokensDTO(accessToken, refreshToken), nil
}
