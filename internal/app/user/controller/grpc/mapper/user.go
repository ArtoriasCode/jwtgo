package mapper

import (
	"jwtgo/internal/app/user/controller/grpc/dto"
	"jwtgo/internal/app/user/entity"
	userPb "jwtgo/internal/generated/proto/user"
)

func MapUserGetByIdRequestToGetByIdRequestDTO(request *userPb.GetByIdRequest) *dto.GetByIdRequestDTO {
	return &dto.GetByIdRequestDTO{
		Id: request.Id,
	}
}

func MapUserDTOToUserGetByIdResponse(dto *dto.UserDTO) *userPb.GetByIdResponse {
	return &userPb.GetByIdResponse{
		User: MapUserDTOToUserProto(dto),
	}
}

func MapUserGetByEmailRequestToGetByEmailRequestDTO(request *userPb.GetByEmailRequest) *dto.GetByEmailRequestDTO {
	return &dto.GetByEmailRequestDTO{
		Email: request.Email,
	}
}

func MapUserDTOToUserGetByEmailResponse(dto *dto.UserDTO) *userPb.GetByEmailResponse {
	return &userPb.GetByEmailResponse{
		User: MapUserDTOToUserProto(dto),
	}
}

func MapUserCreateRequestToCreateRequestDTO(request *userPb.CreateRequest) *dto.CreateRequestDTO {
	return &dto.CreateRequestDTO{
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

func MapUserUpdateRequestToUpdateRequestDTO(request *userPb.UpdateRequest) *dto.UpdateRequestDTO {
	return &dto.UpdateRequestDTO{
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

func MapUserDeleteRequestToDeleteRequestDTO(request *userPb.DeleteRequest) *dto.DeleteRequestDTO {
	return &dto.DeleteRequestDTO{
		Id: request.Id,
	}
}

func MapUserDTOToUserDeleteResponse(dto *dto.UserDTO) *userPb.DeleteResponse {
	return &userPb.DeleteResponse{
		User: MapUserDTOToUserProto(dto),
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

func MapCreateRequestDTOToUserEntity(dto *dto.CreateRequestDTO) *entity.User {
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

func MapUpdateRequestDTOToUserEntity(dto *dto.UpdateRequestDTO) *entity.User {
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
