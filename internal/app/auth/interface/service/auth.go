package service

import (
	"context"

	"jwtgo/internal/app/auth/server/grpc/dto"
)

type AuthService interface {
	SignUp(ctx context.Context, userCredentialsDTO *dto.UserCredentialsDTO) (bool, error)
	SignIn(ctx context.Context, userCredentialsDTO *dto.UserCredentialsDTO) (*dto.UserTokensDTO, error)
	SignOut(ctx context.Context, accessTokenDTO *dto.UserTokenDTO) (bool, error)
	Refresh(ctx context.Context, refreshTokenDTO *dto.UserTokenDTO) (*dto.UserTokensDTO, error)
}
