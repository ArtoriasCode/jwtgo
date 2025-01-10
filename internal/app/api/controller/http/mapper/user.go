package mapper

import (
	"jwtgo/internal/app/api/controller/http/dto"
	pb "jwtgo/internal/proto/auth"
)

func MapUserCredentialsDTOToSignUpRequest(dto *dto.UserCredentialsDTO) *pb.SignUpRequest {
	return &pb.SignUpRequest{
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func MapUserCredentialsDTOToSignInRequest(dto *dto.UserCredentialsDTO) *pb.SignInRequest {
	return &pb.SignInRequest{
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func MapToSignOutRequest(accessToken string) *pb.SignOutRequest {
	return &pb.SignOutRequest{
		AccessToken: accessToken,
	}
}

func MapToRefreshRequest(refreshToken string) *pb.RefreshRequest {
	return &pb.RefreshRequest{
		RefreshToken: refreshToken,
	}
}
