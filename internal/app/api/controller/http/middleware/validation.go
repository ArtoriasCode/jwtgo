package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validator[T any](validate *validator.Validate) gin.HandlerFunc {
	return func(c *gin.Context) {
		var obj T
		if err := c.ShouldBindJSON(&obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request parameters"})
			c.Abort()
			return
		}

		if err := validate.Struct(obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request parameters"})
			c.Abort()
			return
		}

		c.Set("validatedBody", obj)
		c.Next()
	}
}
