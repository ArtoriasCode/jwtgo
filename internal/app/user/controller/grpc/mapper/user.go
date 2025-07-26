package mapper

import (
	"jwtgo/internal/app/user/controller/grpc/dto"
	"jwtgo/internal/app/user/entity"
	userPb "jwtgo/internal/pkg/proto/user"
)

func MapUserGetByIdRequestToUserIdDTO(request *userPb.GetByIdRequest) *dto.UserIdDTO {
	return &dto.UserIdDTO{
		Id: request.Id,
	}
}

func MapUserDTOToUserGetByIdResponse(dto *dto.UserDTO) *userPb.GetByIdResponse {
	return &userPb.GetByIdResponse{
		User: &userPb.User{
			Id:           dto.Id,
			Email:        dto.Email,
			Password:     dto.Password,
			Salt:         dto.Salt,
			RefreshToken: dto.RefreshToken,
			CreatedAt:    dto.CreatedAt,
			UpdatedAt:    dto.CreatedAt,
		},
	}
}

func MapUserGetByEmailRequestToUserEmailDTO(request *userPb.GetByEmailRequest) *dto.UserEmailDTO {
	return &dto.UserEmailDTO{
		Email: request.Email,
	}
}

func MapUserDTOToUserGetByEmailResponse(dto *dto.UserDTO) *userPb.GetByEmailResponse {
	return &userPb.GetByEmailResponse{
		User: &userPb.User{
			Id:           dto.Id,
			Email:        dto.Email,
			Password:     dto.Password,
			Salt:         dto.Salt,
			RefreshToken: dto.RefreshToken,
			CreatedAt:    dto.CreatedAt,
			UpdatedAt:    dto.CreatedAt,
		},
	}
}

func MapUserCreateRequestToUserCreateDTO(request *userPb.CreateRequest) *dto.UserCreateDTO {
	return &dto.UserCreateDTO{
		Email:    request.Email,
		Password: request.Password,
		Salt:     request.Salt,
	}
}

func MapUserDTOToUserCreateResponse(dto *dto.UserDTO) *userPb.CreateResponse {
	return &userPb.CreateResponse{
		User: &userPb.User{
			Id:           dto.Id,
			Email:        dto.Email,
			Password:     dto.Password,
			Salt:         dto.Salt,
			RefreshToken: dto.RefreshToken,
			CreatedAt:    dto.CreatedAt,
			UpdatedAt:    dto.CreatedAt,
		},
	}
}

func MapUserUpdateRequestToUserUpdateDTO(request *userPb.UpdateRequest) *dto.UserUpdateDTO {
	return &dto.UserUpdateDTO{
		Id:           request.Id,
		Email:        request.Email,
		Password:     request.Password,
		Salt:         request.Salt,
		RefreshToken: request.RefreshToken,
	}
}

func MapUserDTOToUserUpdateResponse(dto *dto.UserDTO) *userPb.UpdateResponse {
	return &userPb.UpdateResponse{
		User: &userPb.User{
			Id:           dto.Id,
			Email:        dto.Email,
			Password:     dto.Password,
			Salt:         dto.Salt,
			RefreshToken: dto.RefreshToken,
			CreatedAt:    dto.CreatedAt,
			UpdatedAt:    dto.CreatedAt,
		},
	}
}

func MapUserDeleteRequestToUserIdDTO(request *userPb.DeleteRequest) *dto.UserIdDTO {
	return &dto.UserIdDTO{
		Id: request.Id,
	}
}

func MapStatusToDeleteResponse(status bool) *userPb.DeleteResponse {
	return &userPb.DeleteResponse{
		Success: status,
	}
}

func MapUserEntityToUserDTO(entity *entity.User) *dto.UserDTO {
	return &dto.UserDTO{
		Id:           entity.Id,
		Email:        entity.Email,
		Password:     entity.Password,
		Salt:         entity.Salt,
		RefreshToken: entity.RefreshToken,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.CreatedAt,
	}
}

func MapUserCreateDTOToUserEntity(dto *dto.UserCreateDTO) *entity.User {
	return &entity.User{
		Email:        dto.Email,
		Password:     dto.Password,
		Salt:         dto.Salt,
		RefreshToken: dto.RefreshToken,
	}
}

func MapUserUpdateDTOToUserEntity(dto *dto.UserUpdateDTO) *entity.User {
	return &entity.User{
		Id:           dto.Id,
		Email:        dto.Email,
		Password:     dto.Password,
		Salt:         dto.Salt,
		RefreshToken: dto.RefreshToken,
	}
}
