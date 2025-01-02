package mapper

import (
	"jwtgo/internal/domain/entity"
	"time"

	"jwtgo/internal/controller/http/dto"
)

func MapToUserRefreshTokenDTO(refreshToken string) *dto.UserRefreshTokenDTO {
	return &dto.UserRefreshTokenDTO{
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
