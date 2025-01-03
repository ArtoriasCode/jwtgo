package service

import (
	"jwtgo/internal/app/schema"
)

type TokenService interface {
	GenerateTokens(id string) (string, string, error)
	ValidateToken(signedToken string) (*schema.Claims, error)
}
