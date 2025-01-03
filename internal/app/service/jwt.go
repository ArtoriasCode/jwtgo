package service

import (
	"jwtgo/internal/app/schema"
	customErr "jwtgo/internal/error"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService struct {
	secretKey       string
	accessLifetime  int
	refreshLifetime int
}

func NewJWTService(secretKey string, accessLifetime, refreshLifetime int) *JWTService {
	return &JWTService{
		secretKey:       secretKey,
		accessLifetime:  accessLifetime,
		refreshLifetime: refreshLifetime,
	}
}

func (s *JWTService) GenerateTokens(id string) (string, string, error) {
	accessClaims := &schema.Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(time.Minute * time.Duration(s.accessLifetime)).Unix(),
		},
	}

	refreshClaims := &schema.Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(time.Minute * time.Duration(s.refreshLifetime)).Unix(),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(s.secretKey))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(s.secretKey))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *JWTService) ValidateToken(signedToken string) (*schema.Claims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&schema.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(s.secretKey), nil
		},
	)

	if err != nil {
		return nil, customErr.NewInvalidTokenError("Token is invalid")
	}

	claims, ok := token.Claims.(*schema.Claims)
	if !ok {
		return nil, customErr.NewInvalidTokenError("Token is invalid")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return nil, customErr.NewExpiredTokenError("Token is expired")
	}

	return claims, nil
}
