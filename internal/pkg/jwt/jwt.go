package jwt

import (
	"errors"
	"jwtgo/pkg/logging"
	"time"

	"github.com/golang-jwt/jwt/v5"

	customErr "jwtgo/internal/pkg/error/type"
	"jwtgo/internal/pkg/jwt/schema"
)

type JWTService struct {
	secretKey       string
	accessLifetime  int
	refreshLifetime int
	logger          *logging.Logger
}

func NewJWTService(secretKey string, accessLifetime, refreshLifetime int, logger *logging.Logger) *JWTService {
	return &JWTService{
		secretKey:       secretKey,
		accessLifetime:  accessLifetime,
		refreshLifetime: refreshLifetime,
		logger:          logger,
	}
}

func (s *JWTService) GenerateTokens(id, role, username string) (string, string, customErr.BaseErrorIface) {
	accessClaims := &schema.Claims{
		Id:       id,
		Role:     role,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Minute * time.Duration(s.accessLifetime))),
		},
	}

	refreshClaims := &schema.Claims{
		Id:       id,
		Role:     role,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Minute * time.Duration(s.refreshLifetime))),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(s.secretKey))
	if err != nil {
		s.logger.Error("[JWTService -> GenerateTokens -> NewWithClaims]: ", err)
		return "", "", customErr.NewInternalServerError("Failed to generate access token")
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(s.secretKey))
	if err != nil {
		s.logger.Error("[JWTService -> GenerateTokens -> NewWithClaims]: ", err)
		return "", "", customErr.NewInternalServerError("Failed to generate refresh token")
	}

	return accessToken, refreshToken, nil
}

func (s *JWTService) ValidateToken(signedToken string) (*schema.Claims, customErr.BaseErrorIface) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&schema.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(s.secretKey), nil
		},
	)

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, customErr.NewExpiredTokenError("Token is expired")
		} else {
			return nil, customErr.NewInvalidTokenError("Token is invalid")
		}
	}

	claims, ok := token.Claims.(*schema.Claims)
	if !ok {
		return nil, customErr.NewInvalidTokenError("Token is invalid")
	}

	return claims, nil
}
