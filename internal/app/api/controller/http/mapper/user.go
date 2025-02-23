package mapper

import (
	"jwtgo/internal/app/api/controller/http/dto"
	authPb "jwtgo/internal/pkg/proto/auth"
)

func MapUserCredentialsDTOToAuthSignUpRequest(dto *dto.UserCredentialsDTO) *authPb.SignUpRequest {
	return &authPb.SignUpRequest{
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func MapUserCredentialsDTOToAuthSignInRequest(dto *dto.UserCredentialsDTO) *authPb.SignInRequest {
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
