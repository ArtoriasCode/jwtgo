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
		User: MapUserDTOToUserProto(dto),
	}
}

func MapUserGetByEmailRequestToUserEmailDTO(request *userPb.GetByEmailRequest) *dto.UserEmailDTO {
	return &dto.UserEmailDTO{
		Email: request.Email,
	}
}

func MapUserDTOToUserGetByEmailResponse(dto *dto.UserDTO) *userPb.GetByEmailResponse {
	return &userPb.GetByEmailResponse{
		User: MapUserDTOToUserProto(dto),
	}
}

func MapUserCreateRequestToUserCreateDTO(request *userPb.CreateRequest) *dto.UserCreateDTO {
	return &dto.UserCreateDTO{
		Email:    request.Email,
		Role:     request.Role,
		Username: request.Username,
		Gender:   request.Gender,
		Security: dto.SecurityDTO{
			Password:     request.Security.Password,
			Salt:         request.Security.Salt,
			RefreshToken: request.Security.RefreshToken,
		},
	}
}

func MapUserDTOToUserCreateResponse(dto *dto.UserDTO) *userPb.CreateResponse {
	return &userPb.CreateResponse{
		User: MapUserDTOToUserProto(dto),
	}
}

func MapUserUpdateRequestToUserUpdateDTO(request *userPb.UpdateRequest) *dto.UserUpdateDTO {
	return &dto.UserUpdateDTO{
		Id:       request.Id,
		Email:    request.Email,
		Role:     request.Role,
		Username: request.Username,
		Gender:   request.Gender,
		Security: dto.SecurityDTO{
			Password:     request.Security.Password,
			Salt:         request.Security.Salt,
			RefreshToken: request.Security.RefreshToken,
		},
	}
}

func MapUserDTOToUserUpdateResponse(dto *dto.UserDTO) *userPb.UpdateResponse {
	return &userPb.UpdateResponse{
		User: MapUserDTOToUserProto(dto),
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
		Id:       entity.Id,
		Email:    entity.Email,
		Role:     entity.Role,
		Username: entity.Username,
		Gender:   entity.Gender,
		Security: dto.SecurityDTO{
			Password:     entity.Security.Password,
			Salt:         entity.Security.Salt,
			RefreshToken: entity.Security.RefreshToken,
		},
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func MapUserCreateDTOToUserEntity(dto *dto.UserCreateDTO) *entity.User {
	return &entity.User{
		Email:    dto.Email,
		Role:     dto.Role,
		Username: dto.Username,
		Gender:   dto.Gender,
		Security: entity.Security{
			Password:     dto.Security.Password,
			Salt:         dto.Security.Salt,
			RefreshToken: dto.Security.RefreshToken,
		},
	}
}

func MapUserUpdateDTOToUserEntity(dto *dto.UserUpdateDTO) *entity.User {
	return &entity.User{
		Id:       dto.Id,
		Email:    dto.Email,
		Role:     dto.Role,
		Username: dto.Username,
		Gender:   dto.Gender,
		Security: entity.Security{
			Password:     dto.Security.Password,
			Salt:         dto.Security.Salt,
			RefreshToken: dto.Security.RefreshToken,
		},
	}
}

func MapUserDTOToUserProto(dto *dto.UserDTO) *userPb.User {
	return &userPb.User{
		Id:       dto.Id,
		Email:    dto.Email,
		Role:     dto.Role,
		Username: dto.Username,
		Gender:   dto.Gender,
		Security: &userPb.Security{
			Password:     dto.Security.Password,
			Salt:         dto.Security.Salt,
			RefreshToken: dto.Security.RefreshToken,
		},
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}
