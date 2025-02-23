package service

import (
	"context"

	repositoryInterface "jwtgo/internal/app/user/interface/repository"
	"jwtgo/internal/app/user/server/grpc/dto"
	"jwtgo/internal/app/user/server/grpc/mapper"
	customErr "jwtgo/internal/pkg/error"
	"jwtgo/pkg/logging"
)

type UserService struct {
	userRepository repositoryInterface.UserRepository
	logger         *logging.Logger
}

func NewUserService(
	userRepository repositoryInterface.UserRepository,
	logger *logging.Logger,
) *UserService {
	return &UserService{
		userRepository: userRepository,
		logger:         logger,
	}
}

func (s *UserService) GetById(ctx context.Context, userIdDTO *dto.UserIdDTO) (*dto.UserDTO, error) {
	userEntity, err := s.userRepository.GetById(ctx, userIdDTO.Id)
	if err != nil {
		s.logger.Error("Error while getting user: ", err)
		return nil, customErr.NewInternalServerError("Failed to check user id")
	}

	if userEntity == nil {
		return nil, nil
	}

	return mapper.MapUserEntityToUserDTO(userEntity), nil
}

func (s *UserService) GetByEmail(ctx context.Context, userEmailDTO *dto.UserEmailDTO) (*dto.UserDTO, error) {
	userEntity, err := s.userRepository.GetByEmail(ctx, userEmailDTO.Email)
	if err != nil {
		s.logger.Error("Error while getting user: ", err)
		return nil, customErr.NewInternalServerError("Failed to check user email")
	}

	if userEntity == nil {
		return nil, nil
	}

	return mapper.MapUserEntityToUserDTO(userEntity), nil
}

func (s *UserService) Create(ctx context.Context, userCreateDTO *dto.UserCreateDTO) (*dto.UserDTO, error) {
	userEntity := mapper.MapUserCreateDTOToUserEntity(userCreateDTO)

	createdUserEntity, err := s.userRepository.Create(ctx, userEntity)
	if err != nil {
		s.logger.Error("Error while creating user: ", err)
		return nil, customErr.NewInternalServerError("Failed to create a user")
	}

	return mapper.MapUserEntityToUserDTO(createdUserEntity), nil
}

func (s *UserService) Update(ctx context.Context, userUpdateDTO *dto.UserUpdateDTO) (*dto.UserDTO, error) {
	userEntity := mapper.MapUserUpdateDTOToUserEntity(userUpdateDTO)

	updatedUserEntity, err := s.userRepository.Update(ctx, userEntity.Id, userEntity)
	if err != nil {
		s.logger.Error("Error while updating user: ", err)
		return nil, customErr.NewInternalServerError("Failed to update a user")
	}

	return mapper.MapUserEntityToUserDTO(updatedUserEntity), nil
}

func (s *UserService) Delete(ctx context.Context, userDeleteDTO *dto.UserIdDTO) (bool, error) {
	isUserDeleted, err := s.userRepository.Delete(ctx, userDeleteDTO.Id)
	if err != nil {
		s.logger.Error("Error while deleting user: ", err)
		return false, customErr.NewInternalServerError("Failed to delete a user")
	}

	return isUserDeleted, nil
}
