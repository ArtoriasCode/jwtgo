package repository

import (
	"context"
	customErr "jwtgo/internal/pkg/error/type"

	"jwtgo/internal/app/user/entity"
)

type UserRepository interface {
	GetById(ctx context.Context, id string) (*entity.User, customErr.BaseErrorInterface)
	GetByEmail(ctx context.Context, email string) (*entity.User, customErr.BaseErrorInterface)
	GetAll(ctx context.Context) ([]*entity.User, customErr.BaseErrorInterface)
	Create(ctx context.Context, user *entity.User) (*entity.User, customErr.BaseErrorInterface)
	Update(ctx context.Context, id string, user *entity.User) (*entity.User, customErr.BaseErrorInterface)
	Delete(ctx context.Context, id string) (bool, customErr.BaseErrorInterface)
}
