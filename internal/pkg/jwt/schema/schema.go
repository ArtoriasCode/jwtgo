package schema

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Id   string `json:"sub"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}
