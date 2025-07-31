package mapper

import (
	"jwtgo/internal/app/api/controller/http/dto"
	authPb "jwtgo/internal/pkg/proto/auth"
)

func MapSignUpRequestDTOToAuthSignUpRequest(dto *dto.SignUpRequestDTO) *authPb.SignUpRequest {
	return &authPb.SignUpRequest{
		Email:    dto.Email,
		Password: dto.Password,
		Username: dto.Username,
		Gender:   dto.Gender,
		Role:     "user",
	}
}

func MapSignInRequestDTOToAuthSignInRequest(dto *dto.SignInRequestDTO) *authPb.SignInRequest {
	return &authPb.SignInRequest{
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func MapAccessTokenToAuthSignOutRequest(accessToken string) *authPb.SignOutRequest {
	return &authPb.SignOutRequest{
		AccessToken: accessToken,
	}
}

func MapRefreshTokenToAuthRefreshRequest(refreshToken string) *authPb.RefreshRequest {
	return &authPb.RefreshRequest{
		RefreshToken: refreshToken,
	}
}
