package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"jwtgo/internal/app/api/controller/http/mapper"
	pkgServiceIface "jwtgo/internal/pkg/interface/service"
	authPb "jwtgo/internal/pkg/proto/auth"
	"jwtgo/internal/pkg/request"
	"jwtgo/internal/pkg/request/schema"
)

func Authentication(jwtService pkgServiceIface.JWTServiceIface, authMicroService authPb.AuthServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Request.Cookie("access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
			c.Abort()
			return
		}

		claims, err := jwtService.ValidateToken(accessToken.Value)
		if err != nil {
			refreshToken, err := c.Request.Cookie("refresh_token")
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
				c.Abort()
				return
			}

			ctx := c.Request.Context()
			authRefreshRequest := mapper.MapRefreshTokenToAuthRefreshRequest(refreshToken.Value)

			authRefreshResponse, err := authMicroService.Refresh(ctx, authRefreshRequest)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
				c.Abort()
				return
			}

			request.SetCookies(c, []schema.Cookie{
				{Name: "access_token", Value: authRefreshResponse.AccessToken, Duration: 7 * 24 * time.Hour},
				{Name: "refresh_token", Value: authRefreshResponse.RefreshToken, Duration: 7 * 24 * time.Hour},
			})

			newClaims, err := jwtService.ValidateToken(authRefreshResponse.AccessToken)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "New access token is invalid"})
				c.Abort()
				return
			}

			claims = newClaims
		}

		c.Set("id", claims.Id)
		c.Set("role", claims.Role)
		c.Next()
	}
}
