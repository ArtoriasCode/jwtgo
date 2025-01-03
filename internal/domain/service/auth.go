package service

import (
	"context"
	"jwtgo/internal/controller/http/dto"
	"jwtgo/internal/controller/http/mapper"
	domainEntity "jwtgo/internal/domain/entity"
	customErr "jwtgo/internal/error"
	"jwtgo/pkg/logging"
	"jwtgo/pkg/security"
	"time"
)

type UserRepository interface {
	GetById(ctx context.Context, id string) (*domainEntity.User, error)
	GetByEmail(ctx context.Context, email string) (*domainEntity.User, error)
	GetAll(ctx context.Context) ([]*domainEntity.User, error)
	Create(ctx context.Context, domainUser *domainEntity.User) (bool, error)
	Update(ctx context.Context, id string, domainUser *domainEntity.User) (bool, error)
	Delete(ctx context.Context, id string) (bool, error)
}

type UserService struct {
	repository      UserRepository
	tokenManager    *security.TokenManager
	passwordManager *security.PasswordManager
	logger          *logging.Logger
}

func NewUserService(
	repository UserRepository,
	tokenManager *security.TokenManager,
	passwordManager *security.PasswordManager,
	logger *logging.Logger,
) *UserService {
	return &UserService{
		repository:      repository,
		tokenManager:    tokenManager,
		passwordManager: passwordManager,
		logger:          logger,
	}
}

func (s *UserService) SignUp(ctx context.Context, userCredentialsDTO *dto.UserCredentialsDTO) (bool, error) {
	existingUserEntity, err := s.repository.GetByEmail(ctx, userCredentialsDTO.Email)
	if err != nil {
		s.logger.Error("Error while getting user: ", err)
		return false, customErr.NewInternalServerError("Failed to check user email")
	}

	if existingUserEntity != nil {
		return false, customErr.NewAlreadyExistsError("Email already exists")
	}

	localSalt, err := s.passwordManager.GenerateSalt(32)
	if err != nil {
		s.logger.Error("Error while generating local salt: ", err)
		return false, customErr.NewInternalServerError("Failed to create a user")
	}

	hashedPassword, err := s.passwordManager.HashPassword(userCredentialsDTO.Password, localSalt)
	if err != nil {
		s.logger.Error("Error while hashing password: ", err)
		return false, customErr.NewInternalServerError("Failed to create a user")
	}

	userCredentialsDTO.Password = hashedPassword

	userCreateEntity := mapper.MapUserCredentialsDTOToDomainUser(userCredentialsDTO)
	userCreateEntity.Salt = localSalt

	_, err = s.repository.Create(ctx, userCreateEntity)
	if err != nil {
		s.logger.Error("Error while creating user: ", err)
		return false, customErr.NewInternalServerError("Failed to create a user")
	}

	return true, nil
}

func (s *UserService) SignIn(ctx context.Context, userCredentialsDTO *dto.UserCredentialsDTO) (*dto.UserTokensDTO, error) {
	existingUserEntity, err := s.repository.GetByEmail(ctx, userCredentialsDTO.Email)
	if err != nil {
		s.logger.Error("Error while getting user: ", err)
		return nil, customErr.NewInternalServerError("Failed to check user email")
	}

	if existingUserEntity == nil {
		return nil, customErr.NewInvalidCredentialsError("Invalid login or password")
	}

	passwordIsValid := s.passwordManager.VerifyPassword(userCredentialsDTO.Password, existingUserEntity.Password, existingUserEntity.Salt)
	if !passwordIsValid {
		return nil, customErr.NewInvalidCredentialsError("Invalid login or password")
	}

	accessToken, refreshToken, err := s.tokenManager.GenerateTokens(existingUserEntity.Id)
	if err != nil {
		s.logger.Error("Error while generating tokens: ", err)
		return nil, customErr.NewInternalServerError("Token generation error")
	}

	existingUserEntity.RefreshToken = refreshToken
	existingUserEntity.UpdatedAt = time.Now().UTC()

	_, err = s.repository.Update(ctx, existingUserEntity.Id, existingUserEntity)
	if err != nil {
		s.logger.Error("Error while updating user: ", err)
		return nil, customErr.NewInternalServerError("Token updating error")
	}

	return mapper.MapToUserTokensDTO(accessToken, refreshToken), nil
}

func (s *UserService) Refresh(ctx context.Context, refreshTokenDTO *dto.UserRefreshTokenDTO) (*dto.UserTokensDTO, error) {
	claims, err := s.tokenManager.ValidateToken(refreshTokenDTO.RefreshToken)
	if err != nil {
		s.logger.Error("Error while checking refresh token: ", err)
		return nil, err
	}

	existingUserEntity, err := s.repository.GetById(ctx, claims.Id)
	if err != nil {
		s.logger.Error("Error while getting user: ", err)
		return nil, customErr.NewInternalServerError("Failed to check user id")
	}

	if existingUserEntity == nil {
		return nil, customErr.NewUserNotFoundError("User not found")
	}

	if refreshTokenDTO.RefreshToken != existingUserEntity.RefreshToken {
		return nil, customErr.NewInvalidRefreshTokenError("Invalid refresh token")
	}

	accessToken, refreshToken, err := s.tokenManager.GenerateTokens(existingUserEntity.Id)
	if err != nil {
		s.logger.Error("Error while generating tokens: ", err)
		return nil, customErr.NewInternalServerError("Token generation error")
	}

	existingUserEntity.RefreshToken = refreshToken
	existingUserEntity.UpdatedAt = time.Now().UTC()

	_, err = s.repository.Update(ctx, claims.Id, existingUserEntity)
	if err != nil {
		s.logger.Error("Error while updating user: ", err)
		return nil, customErr.NewInternalServerError("Token updating error")
	}

	return mapper.MapToUserTokensDTO(accessToken, refreshToken), nil
}
