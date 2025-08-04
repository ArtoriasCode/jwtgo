package mapper

import (
	"jwtgo/internal/app/api/controller/http/dto"
	authPb "jwtgo/internal/generated/proto/auth"
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

func MapAuthSignUpResponseToSignUpResponseDTO(response *authPb.SignUpResponse) *dto.SignUpResponseDTO {
	return &dto.SignUpResponseDTO{
		Message: response.Message,
	}
}

func MapSignInRequestDTOToAuthSignInRequest(dto *dto.SignInRequestDTO) *authPb.SignInRequest {
	return &authPb.SignInRequest{
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func MapAuthSignInResponseToSignInResponseDTO(response *authPb.SignInResponse) *dto.SignInResponseDTO {
	return &dto.SignInResponseDTO{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
		Message:      response.Message,
	}
}

func MapUserIdToSignOutRequestDTO(id string) *dto.SignOutRequestDTO {
	return &dto.SignOutRequestDTO{
		Id: id,
	}
}

func MapSignOutRequestDTOToAuthSignOutRequest(dto *dto.SignOutRequestDTO) *authPb.SignOutRequest {
	return &authPb.SignOutRequest{
		Id: dto.Id,
	}
}

func MapAuthSignOutResponseToSignOutResponseDTO(response *authPb.SignOutResponse) *dto.SignOutResponseDTO {
	return &dto.SignOutResponseDTO{
		Message: response.Message,
	}
}

func MapRefreshTokenToRefreshRequestDTO(refreshToken string) *dto.RefreshRequestDTO {
	return &dto.RefreshRequestDTO{
		RefreshToken: refreshToken,
	}
}

func MapRefreshRequestDTOToAuthRefreshRequest(dto *dto.RefreshRequestDTO) *authPb.RefreshRequest {
	return &authPb.RefreshRequest{
		RefreshToken: dto.RefreshToken,
	}
}

func MapAuthRefreshResponseToRefreshResponseDTO(response *authPb.RefreshResponse) *dto.RefreshResponseDTO {
	return &dto.RefreshResponseDTO{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
		Message:      response.Message,
	}
}
