package mapper

import (
	"jwtgo/internal/app/auth/controller/grpc/dto"
	authPb "jwtgo/internal/pkg/proto/auth"
	userPb "jwtgo/internal/pkg/proto/user"
)

func MapTokensToUserTokensDTO(accessToken, refreshToken string) *dto.UserTokensDTO {
	return &dto.UserTokensDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func MapAuthSignUpRequestToUserCredentialsDTO(request *authPb.SignUpRequest) *dto.UserCredentialsDTO {
	return &dto.UserCredentialsDTO{
		Email:    request.Email,
		Password: request.Password,
		Role:     request.Role,
	}
}

func MapAuthSignInRequestToUserCredentialsDTO(request *authPb.SignInRequest) *dto.UserCredentialsDTO {
	return &dto.UserCredentialsDTO{
		Email:    request.Email,
		Password: request.Password,
	}
}

func MapUserTokensDTOToAuthSignInResponse(dto *dto.UserTokensDTO, message string) *authPb.SignInResponse {
	return &authPb.SignInResponse{
		AccessToken:  dto.AccessToken,
		RefreshToken: dto.RefreshToken,
		Message:      message,
	}
}

func MapAuthSignOutRequestToUserTokenDTO(request *authPb.SignOutRequest) *dto.UserTokenDTO {
	return &dto.UserTokenDTO{
		Token: request.AccessToken,
	}
}

func MapAuthRefreshRequestToUserTokenDTO(request *authPb.RefreshRequest) *dto.UserTokenDTO {
	return &dto.UserTokenDTO{
		Token: request.RefreshToken,
	}
}

func MapUserTokensDTOToAuthRefreshResponse(dto *dto.UserTokensDTO, message string) *authPb.RefreshResponse {
	return &authPb.RefreshResponse{
		AccessToken:  dto.AccessToken,
		RefreshToken: dto.RefreshToken,
		Message:      message,
	}
}

func MapIdToUserGetByIdRequest(id string) *userPb.GetByIdRequest {
	return &userPb.GetByIdRequest{
		Id: id,
	}
}

func MapEmailToUserGetByEmailRequest(email string) *userPb.GetByEmailRequest {
	return &userPb.GetByEmailRequest{
		Email: email,
	}
}

func MapUserCredentialsDTOToUserCreateRequest(dto *dto.UserCredentialsDTO) *userPb.CreateRequest {
	return &userPb.CreateRequest{
		Email: dto.Email,
		Role:  dto.Role,
		Security: &userPb.Security{
			Password: dto.Password,
		},
	}
}

func MapUserGetByEmailResponseToUserUpdateRequest(response *userPb.GetByEmailResponse) *userPb.UpdateRequest {
	return &userPb.UpdateRequest{
		Id:    response.User.Id,
		Email: response.User.Email,
		Role:  response.User.Role,
		Security: &userPb.Security{
			Password:     response.User.Security.Password,
			Salt:         response.User.Security.Salt,
			RefreshToken: response.User.Security.RefreshToken,
		},
	}
}

func MapUserGetByIdResponseToUserUpdateRequest(response *userPb.GetByIdResponse) *userPb.UpdateRequest {
	return &userPb.UpdateRequest{
		Id:    response.User.Id,
		Email: response.User.Email,
		Role:  response.User.Role,
		Security: &userPb.Security{
			Password:     response.User.Security.Password,
			Salt:         response.User.Security.Salt,
			RefreshToken: response.User.Security.RefreshToken,
		},
	}
}
