package service

import (
	"jwtgo/internal/app/schema"
)

type JWTService interface {
	GenerateTokens(id string) (string, string, error)
	ValidateToken(signedToken string) (*schema.Claims, error)
}
