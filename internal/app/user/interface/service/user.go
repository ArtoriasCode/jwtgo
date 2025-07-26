package service

import (
	"context"

	"jwtgo/internal/app/user/controller/grpc/dto"
	customErr "jwtgo/internal/pkg/error/type"
)

type UserServiceIface interface {
	GetById(ctx context.Context, userIdDTO *dto.UserIdDTO) (*dto.UserDTO, customErr.BaseErrorIface)
	GetByEmail(ctx context.Context, userEmailDTO *dto.UserEmailDTO) (*dto.UserDTO, customErr.BaseErrorIface)
	Create(ctx context.Context, userCreateDTO *dto.UserCreateDTO) (*dto.UserDTO, customErr.BaseErrorIface)
	Update(ctx context.Context, useDTO *dto.UserUpdateDTO) (*dto.UserDTO, customErr.BaseErrorIface)
	Delete(ctx context.Context, userIdDTO *dto.UserIdDTO) (bool, customErr.BaseErrorIface)
}
