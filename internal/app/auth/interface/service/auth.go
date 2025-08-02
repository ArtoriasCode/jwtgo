package service

import (
	"context"

	"jwtgo/internal/app/auth/controller/grpc/dto"
	customErr "jwtgo/internal/pkg/error/type"
)

type AuthServiceIface interface {
	SignUp(ctx context.Context, signUpRequestDTO *dto.SignUpRequestDTO) (*dto.SignUpResponseDTO, customErr.BaseErrorIface)
	SignIn(ctx context.Context, signInRequestDTO *dto.SignInRequestDTO) (*dto.SignInResponseDTO, customErr.BaseErrorIface)
	SignOut(ctx context.Context, signOutRequestDTO *dto.SignOutRequestDTO) (*dto.SignOutResponseDTO, customErr.BaseErrorIface)
	Refresh(ctx context.Context, refreshRequestDTO *dto.RefreshRequestDTO) (*dto.RefreshResponseDTO, customErr.BaseErrorIface)
}
