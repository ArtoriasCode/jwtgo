package service

import (
	"context"

	"jwtgo/internal/app/auth/controller/grpc/dto"
	customErr "jwtgo/internal/pkg/error/type"
)

type AuthServiceIface interface {
	SignUp(ctx context.Context, signUpRequestDTO *dto.SignUpRequestDTO) (bool, customErr.BaseErrorIface)
	SignIn(ctx context.Context, signInRequestDTO *dto.SignInRequestDTO) (*dto.UserTokensDTO, customErr.BaseErrorIface)
	SignOut(ctx context.Context, signOutRequestDTO *dto.SignOutRequestDTO) (bool, customErr.BaseErrorIface)
	Refresh(ctx context.Context, refreshTokenDTO *dto.UserTokenDTO) (*dto.UserTokensDTO, customErr.BaseErrorIface)
}
