package schema

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Id string `json:"sub"`
	jwt.StandardClaims
}
