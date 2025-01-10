package service

import (
	"jwtgo/internal/pkg/service/schema"
)

type JWTService interface {
	GenerateTokens(id string) (string, string, error)
	ValidateToken(signedToken string) (*schema.Claims, error)
}
