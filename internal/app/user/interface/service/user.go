package service

import (
	"context"

	"jwtgo/internal/app/user/controller/grpc/dto"
	customErr "jwtgo/internal/pkg/error/type"
)

type UserServiceIface interface {
	GetById(ctx context.Context, getByIdRequestDTO *dto.GetByIdRequestDTO) (*dto.UserDTO, customErr.BaseErrorIface)
	GetByEmail(ctx context.Context, getByEmailRequestDTO *dto.GetByEmailRequestDTO) (*dto.UserDTO, customErr.BaseErrorIface)
	Create(ctx context.Context, createRequestDTO *dto.CreateRequestDTO) (*dto.UserDTO, customErr.BaseErrorIface)
	Update(ctx context.Context, updateRequestDTO *dto.UpdateRequestDTO) (*dto.UserDTO, customErr.BaseErrorIface)
	Delete(ctx context.Context, deleteRequestDTO *dto.DeleteRequestDTO) (*dto.UserDTO, customErr.BaseErrorIface)
}
