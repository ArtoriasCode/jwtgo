package repository

import (
	"context"

	"jwtgo/internal/app/user/entity"
	customErr "jwtgo/internal/pkg/error/type"
)

type UserRepositoryIface interface {
	GetById(ctx context.Context, id string) (*entity.User, customErr.BaseErrorIface)
	GetByEmail(ctx context.Context, email string) (*entity.User, customErr.BaseErrorIface)
	GetAll(ctx context.Context) ([]*entity.User, customErr.BaseErrorIface)
	Create(ctx context.Context, user *entity.User) (*entity.User, customErr.BaseErrorIface)
	Update(ctx context.Context, id string, user *entity.User) (*entity.User, customErr.BaseErrorIface)
	Delete(ctx context.Context, id string) (*entity.User, customErr.BaseErrorIface)
}
