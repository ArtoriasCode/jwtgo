package service

import (
	"context"
	"jwtgo/internal/controller/http/dto"
)

type AuthService interface {
	SignUp(ctx context.Context, userCredentialsDTO *dto.UserCredentialsDTO) (bool, error)
	SignIn(ctx context.Context, userCredentialsDTO *dto.UserCredentialsDTO) (*dto.UserTokensDTO, error)
	Refresh(ctx context.Context, refreshTokenDTO *dto.UserRefreshTokenDTO) (*dto.UserTokensDTO, error)
}
