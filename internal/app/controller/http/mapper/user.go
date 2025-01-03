package mapper

import (
	"jwtgo/internal/app/controller/http/dto"
	"jwtgo/internal/app/entity"
	"time"
)

func MapToUserRefreshTokenDTO(refreshToken string) *dto.UserRefreshTokenDTO {
	return &dto.UserRefreshTokenDTO{
		RefreshToken: refreshToken,
	}
}

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
