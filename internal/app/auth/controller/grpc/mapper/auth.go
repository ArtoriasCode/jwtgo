package mapper

import (
	"jwtgo/internal/app/auth/controller/grpc/dto"
	authPb "jwtgo/internal/pkg/proto/auth"
	userPb "jwtgo/internal/pkg/proto/user"
)

func MapAuthSignUpRequestToSignUpRequestDTO(request *authPb.SignUpRequest) *dto.SignUpRequestDTO {
	return &dto.SignUpRequestDTO{
		Email:    request.Email,
		Password: request.Password,
		Role:     request.Role,
		Username: request.Username,
		Gender:   request.Gender,
	}
}

func MapSignUpRequestDTOToUserCreateRequest(dto *dto.SignUpRequestDTO) *userPb.CreateRequest {
	return &userPb.CreateRequest{
		Email:    dto.Email,
		Role:     dto.Role,
		Username: dto.Username,
		Gender:   dto.Gender,
		Security: &userPb.Security{
			Password: dto.Password,
		},
	}
}

func MapUserCreateResponseToAuthSignUpResponseDTO(response *userPb.CreateResponse) *dto.SignUpResponseDTO {
	return &dto.SignUpResponseDTO{
		Email:    response.User.Email,
		Role:     response.User.Role,
		Username: response.User.Username,
		Gender:   response.User.Gender,
		Security: dto.SecurityDTO{
			Password:     response.User.Security.Password,
			Salt:         response.User.Security.Salt,
			RefreshToken: response.User.Security.RefreshToken,
		},
	}
}

func MapAuthSignInRequestToSignInRequestDTO(request *authPb.SignInRequest) *dto.SignInRequestDTO {
	return &dto.SignInRequestDTO{
		Email:    request.Email,
		Password: request.Password,
	}
}

func MapTokensToSignInResponseDTO(accessToken, refreshToken string) *dto.SignInResponseDTO {
	return &dto.SignInResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func MapSignInResponseDTOToAuthSignInResponse(dto *dto.SignInResponseDTO, message string) *authPb.SignInResponse {
	return &authPb.SignInResponse{
		AccessToken:  dto.AccessToken,
		RefreshToken: dto.RefreshToken,
		Message:      message,
	}
}

func MapAuthSignOutRequestToSignOutRequestDTO(request *authPb.SignOutRequest) *dto.SignOutRequestDTO {
	return &dto.SignOutRequestDTO{
		Id: request.Id,
	}
}

func MapIsSignedOutToAuthSignOutResponseDTO(status bool) *dto.SignOutResponseDTO {
	return &dto.SignOutResponseDTO{
		IsSignedOut: status,
	}
}

func MapAuthRefreshRequestToRefreshRequestDTO(request *authPb.RefreshRequest) *dto.RefreshRequestDTO {
	return &dto.RefreshRequestDTO{
		RefreshToken: request.RefreshToken,
	}
}

func MapTokensToRefreshResponseDTO(accessToken, refreshToken string) *dto.RefreshResponseDTO {
	return &dto.RefreshResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func MapRefreshResponseDTOToAuthRefreshResponse(dto *dto.RefreshResponseDTO, message string) *authPb.RefreshResponse {
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

func MapUserGetByEmailResponseToUserUpdateRequest(response *userPb.GetByEmailResponse) *userPb.UpdateRequest {
	return &userPb.UpdateRequest{
		Id:       response.User.Id,
		Email:    response.User.Email,
		Role:     response.User.Role,
		Username: response.User.Username,
		Gender:   response.User.Gender,
		Security: &userPb.Security{
			Password:     response.User.Security.Password,
			Salt:         response.User.Security.Salt,
			RefreshToken: response.User.Security.RefreshToken,
		},
	}
}

func MapUserGetByIdResponseToUserUpdateRequest(response *userPb.GetByIdResponse) *userPb.UpdateRequest {
	return &userPb.UpdateRequest{
		Id:       response.User.Id,
		Email:    response.User.Email,
		Role:     response.User.Role,
		Username: response.User.Username,
		Gender:   response.User.Gender,
		Security: &userPb.Security{
			Password:     response.User.Security.Password,
			Salt:         response.User.Security.Salt,
			RefreshToken: response.User.Security.RefreshToken,
		},
	}
}
