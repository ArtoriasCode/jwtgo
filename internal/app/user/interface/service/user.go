package service

import (
	"context"
	customErr "jwtgo/internal/pkg/error/type"

	"jwtgo/internal/app/user/controller/grpc/dto"
)

type UserService interface {
	GetById(ctx context.Context, userIdDTO *dto.UserIdDTO) (*dto.UserDTO, customErr.BaseErrorInterface)
	GetByEmail(ctx context.Context, userEmailDTO *dto.UserEmailDTO) (*dto.UserDTO, customErr.BaseErrorInterface)
	Create(ctx context.Context, userCreateDTO *dto.UserCreateDTO) (*dto.UserDTO, customErr.BaseErrorInterface)
	Update(ctx context.Context, useDTO *dto.UserUpdateDTO) (*dto.UserDTO, customErr.BaseErrorInterface)
	Delete(ctx context.Context, userIdDTO *dto.UserIdDTO) (bool, customErr.BaseErrorInterface)
}
