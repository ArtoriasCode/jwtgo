package mapper

import (
	"jwtgo/internal/app/api/controller/http/dto"
	authPb "jwtgo/internal/pkg/proto/auth"
)

func MapUserSignUpDTOToAuthSignUpRequest(dto *dto.UserSignUpDTO) *authPb.SignUpRequest {
	return &authPb.SignUpRequest{
		Email:    dto.Email,
		Password: dto.Password,
		Username: dto.Username,
		Gender:   dto.Gender,
		Role:     "user",
	}
}

func MapUserSignInDTOToAuthSignInRequest(dto *dto.UserSignInDTO) *authPb.SignInRequest {
	return &authPb.SignInRequest{
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func MapAccessTokenToAuthRefreshRequest(accessToken string) *authPb.SignOutRequest {
	return &authPb.SignOutRequest{
		AccessToken: accessToken,
	}
}

func MapRefreshTokenToAuthRefreshRequest(refreshToken string) *authPb.RefreshRequest {
	return &authPb.RefreshRequest{
		RefreshToken: refreshToken,
	}
}
