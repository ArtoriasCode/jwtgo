package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	clientInterface "jwtgo/internal/app/interface/service"
)

func Authentication(jwtService clientInterface.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Request.Cookie("access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
			c.Abort()
			return
		}

		claims, err := jwtService.ValidateToken(accessToken.Value)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("id", claims.Id)
		c.Next()
	}
}
