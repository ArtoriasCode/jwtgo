package service

import (
	"context"
	"errors"

	"jwtgo/internal/app/user/controller/grpc/dto"
	"jwtgo/internal/app/user/controller/grpc/mapper"
	repositoryInterface "jwtgo/internal/app/user/interface/repository"
	customErr "jwtgo/internal/pkg/error/type"
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

func (s *UserService) GetById(ctx context.Context, userIdDTO *dto.UserIdDTO) (*dto.UserDTO, customErr.BaseErrorInterface) {
	userEntity, err := s.userRepository.GetById(ctx, userIdDTO.Id)
	if err != nil {
		var notFoundErr *customErr.NotFoundError
		if errors.As(err, &notFoundErr) {
			return nil, nil
		}
		return nil, err
	}

	return mapper.MapUserEntityToUserDTO(userEntity), nil
}

func (s *UserService) GetByEmail(ctx context.Context, userEmailDTO *dto.UserEmailDTO) (*dto.UserDTO, customErr.BaseErrorInterface) {
	userEntity, err := s.userRepository.GetByEmail(ctx, userEmailDTO.Email)
	if err != nil {
		var notFoundErr *customErr.NotFoundError
		if errors.As(err, &notFoundErr) {
			return nil, nil
		}
		return nil, err
	}

	return mapper.MapUserEntityToUserDTO(userEntity), nil
}

func (s *UserService) Create(ctx context.Context, userCreateDTO *dto.UserCreateDTO) (*dto.UserDTO, customErr.BaseErrorInterface) {
	userEntity := mapper.MapUserCreateDTOToUserEntity(userCreateDTO)

	createdUserEntity, err := s.userRepository.Create(ctx, userEntity)
	if err != nil {
		return nil, err
	}

	return mapper.MapUserEntityToUserDTO(createdUserEntity), nil
}

func (s *UserService) Update(ctx context.Context, userUpdateDTO *dto.UserUpdateDTO) (*dto.UserDTO, customErr.BaseErrorInterface) {
	userEntity := mapper.MapUserUpdateDTOToUserEntity(userUpdateDTO)

	updatedUserEntity, err := s.userRepository.Update(ctx, userEntity.Id, userEntity)
	if err != nil {
		return nil, err
	}

	return mapper.MapUserEntityToUserDTO(updatedUserEntity), nil
}

func (s *UserService) Delete(ctx context.Context, userDeleteDTO *dto.UserIdDTO) (bool, customErr.BaseErrorInterface) {
	isUserDeleted, err := s.userRepository.Delete(ctx, userDeleteDTO.Id)
	if err != nil {
		return false, err
	}

	return isUserDeleted, nil
}
