package repository

import (
	"context"

	"jwtgo/internal/app/user/entity"
)

type UserRepository interface {
	GetById(ctx context.Context, id string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetAll(ctx context.Context) ([]*entity.User, error)
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	Update(ctx context.Context, id string, user *entity.User) (*entity.User, error)
	Delete(ctx context.Context, id string) (bool, error)
}
