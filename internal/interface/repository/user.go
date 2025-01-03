package repository

import (
	"context"
	domainEntity "jwtgo/internal/domain/entity"
)

type UserRepository interface {
	GetById(ctx context.Context, id string) (*domainEntity.User, error)
	GetByEmail(ctx context.Context, email string) (*domainEntity.User, error)
	GetAll(ctx context.Context) ([]*domainEntity.User, error)
	Create(ctx context.Context, domainUser *domainEntity.User) (bool, error)
	Update(ctx context.Context, id string, domainUser *domainEntity.User) (bool, error)
	Delete(ctx context.Context, id string) (bool, error)
}
