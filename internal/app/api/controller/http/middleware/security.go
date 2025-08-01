package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	pkgServiceIface "jwtgo/internal/pkg/interface/service"
)

func Authentication(jwtService pkgServiceIface.JWTServiceIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access_token")
		if err != nil || accessToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
			c.Abort()
			return
		}

		claims, err := jwtService.ValidateToken(accessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("id", claims.Id)
		c.Set("role", claims.Role)
		c.Set("username", claims.Username)
		c.Next()
	}
}
