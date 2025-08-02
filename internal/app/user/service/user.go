package service

import (
	"context"
	"errors"

	"jwtgo/internal/app/user/controller/grpc/dto"
	"jwtgo/internal/app/user/controller/grpc/mapper"
	userRepositoryIface "jwtgo/internal/app/user/interface/repository"
	customErr "jwtgo/internal/pkg/error/type"
	"jwtgo/pkg/logging"
)

type UserService struct {
	userRepository userRepositoryIface.UserRepositoryIface
	logger         *logging.Logger
}

func NewUserService(
	userRepository userRepositoryIface.UserRepositoryIface,
	logger *logging.Logger,
) *UserService {
	return &UserService{
		userRepository: userRepository,
		logger:         logger,
	}
}

func (s *UserService) handleError(err error) (*dto.UserDTO, customErr.BaseErrorIface) {
	var notFoundErr *customErr.NotFoundError
	if errors.As(err, &notFoundErr) {
		return nil, nil
	}

	return nil, err
}

func (s *UserService) GetById(ctx context.Context, getByIdRequestDTO *dto.GetByIdRequestDTO) (*dto.UserDTO, customErr.BaseErrorIface) {
	foundUserEntity, err := s.userRepository.GetById(ctx, getByIdRequestDTO.Id)
	if err != nil {
		return s.handleError(err)
	}

	return mapper.MapUserEntityToUserDTO(foundUserEntity), nil
}

func (s *UserService) GetByEmail(ctx context.Context, getByEmailRequestDTO *dto.GetByEmailRequestDTO) (*dto.UserDTO, customErr.BaseErrorIface) {
	foundUserEntity, err := s.userRepository.GetByEmail(ctx, getByEmailRequestDTO.Email)
	if err != nil {
		return s.handleError(err)
	}

	return mapper.MapUserEntityToUserDTO(foundUserEntity), nil
}

func (s *UserService) Create(ctx context.Context, createRequestDTO *dto.CreateRequestDTO) (*dto.UserDTO, customErr.BaseErrorIface) {
	userEntity := mapper.MapCreateRequestDTOToUserEntity(createRequestDTO)

	createdUserEntity, err := s.userRepository.Create(ctx, userEntity)
	if err != nil {
		return s.handleError(err)
	}

	return mapper.MapUserEntityToUserDTO(createdUserEntity), nil
}

func (s *UserService) Update(ctx context.Context, updateRequestDTO *dto.UpdateRequestDTO) (*dto.UserDTO, customErr.BaseErrorIface) {
	userEntity := mapper.MapUpdateRequestDTOToUserEntity(updateRequestDTO)

	updatedUserEntity, err := s.userRepository.Update(ctx, userEntity.Id, userEntity)
	if err != nil {
		return s.handleError(err)
	}

	return mapper.MapUserEntityToUserDTO(updatedUserEntity), nil
}

func (s *UserService) Delete(ctx context.Context, deleteRequestDTO *dto.DeleteRequestDTO) (*dto.UserDTO, customErr.BaseErrorIface) {
	deletedUserEntity, err := s.userRepository.Delete(ctx, deleteRequestDTO.Id)
	if err != nil {
		return s.handleError(err)
	}

	return mapper.MapUserEntityToUserDTO(deletedUserEntity), nil
}
