package service

import (
	customErr "jwtgo/internal/pkg/error/type"
	"jwtgo/internal/pkg/service/schema"
)

type JWTService interface {
	GenerateTokens(id string) (string, string, customErr.BaseErrorInterface)
	ValidateToken(signedToken string) (*schema.Claims, customErr.BaseErrorInterface)
}
