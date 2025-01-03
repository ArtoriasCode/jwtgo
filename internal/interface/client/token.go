package client

import (
	"jwtgo/pkg/security/schema"
)

type TokenManager interface {
	GenerateTokens(id string) (string, string, error)
	ValidateToken(signedToken string) (*schema.Claims, error)
}
