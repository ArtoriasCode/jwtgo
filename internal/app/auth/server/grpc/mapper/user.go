package mapper

import (
	"time"

	"jwtgo/internal/app/auth/entity"
	"jwtgo/internal/app/auth/server/grpc/dto"
	pb "jwtgo/internal/proto/auth"
)

func MapToUserTokensDTO(accessToken, refreshToken string) *dto.UserTokensDTO {
	return &dto.UserTokensDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func MapUserCredentialsDTOToDomainUser(userCredentialsDTO *dto.UserCredentialsDTO) *entity.User {
	now := time.Now().UTC()

	return &entity.User{
		Email:     userCredentialsDTO.Email,
		Password:  userCredentialsDTO.Password,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func MapSignUpRequestToUserCredentialsDTO(request *pb.SignUpRequest) *dto.UserCredentialsDTO {
	return &dto.UserCredentialsDTO{
		Email:    request.Email,
		Password: request.Password,
	}
}

func MapSignInRequestToUserCredentialsDTO(request *pb.SignInRequest) *dto.UserCredentialsDTO {
	return &dto.UserCredentialsDTO{
		Email:    request.Email,
		Password: request.Password,
	}
}

func MapUserTokensDTOToSignInResponse(dto *dto.UserTokensDTO, message string) *pb.SignInResponse {
	return &pb.SignInResponse{
		AccessToken:  dto.AccessToken,
		RefreshToken: dto.RefreshToken,
		Message:      message,
	}
}

func MapSignOutRequestToUserTokenDTO(request *pb.SignOutRequest) *dto.UserTokenDTO {
	return &dto.UserTokenDTO{
		Token: request.AccessToken,
	}
}

func MapRefreshRequestToUserTokenDTO(request *pb.RefreshRequest) *dto.UserTokenDTO {
	return &dto.UserTokenDTO{
		Token: request.RefreshToken,
	}
}

func MapUserTokensDTOToRefreshResponse(dto *dto.UserTokensDTO, message string) *pb.RefreshResponse {
	return &pb.RefreshResponse{
		AccessToken:  dto.AccessToken,
		RefreshToken: dto.RefreshToken,
		Message:      message,
	}
}
