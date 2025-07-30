package service

import (
	customErr "jwtgo/internal/pkg/error/type"
	"jwtgo/internal/pkg/jwt/schema"
)

type JWTServiceIface interface {
	GenerateTokens(id, role string) (string, string, customErr.BaseErrorIface)
	ValidateToken(signedToken string) (*schema.Claims, customErr.BaseErrorIface)
}
