package service

import (
	"context"
	customErr "jwtgo/internal/pkg/error/type"

	"jwtgo/internal/app/auth/controller/grpc/dto"
)

type AuthService interface {
	SignUp(ctx context.Context, userCredentialsDTO *dto.UserCredentialsDTO) (bool, customErr.BaseErrorInterface)
	SignIn(ctx context.Context, userCredentialsDTO *dto.UserCredentialsDTO) (*dto.UserTokensDTO, customErr.BaseErrorInterface)
	SignOut(ctx context.Context, accessTokenDTO *dto.UserTokenDTO) (bool, customErr.BaseErrorInterface)
	Refresh(ctx context.Context, refreshTokenDTO *dto.UserTokenDTO) (*dto.UserTokensDTO, customErr.BaseErrorInterface)
}
