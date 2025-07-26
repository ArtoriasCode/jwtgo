package service

import (
	"context"

	"jwtgo/internal/app/user/controller/grpc/dto"
)

type UserService interface {
	GetById(ctx context.Context, userIdDTO *dto.UserIdDTO) (*dto.UserDTO, error)
	GetByEmail(ctx context.Context, userEmailDTO *dto.UserEmailDTO) (*dto.UserDTO, error)
	Create(ctx context.Context, userCreateDTO *dto.UserCreateDTO) (*dto.UserDTO, error)
	Update(ctx context.Context, useDTO *dto.UserUpdateDTO) (*dto.UserDTO, error)
	Delete(ctx context.Context, userIdDTO *dto.UserIdDTO) (bool, error)
}
