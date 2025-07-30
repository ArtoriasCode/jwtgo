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
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
			c.Abort()
			return
		}

		if err := validate.Struct(obj); err != nil {
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				errors := make(map[string]string)
				for _, e := range validationErrors {
					fieldName := e.Field()
					tag := e.Tag()
					param := e.Param()

					var msg string
					switch tag {
					case "required":
						msg = fieldName + " is required"
					case "email":
						msg = fieldName + " must be a valid email address"
					case "min":
						msg = fieldName + " must be at least " + param + " characters long"
					case "max":
						msg = fieldName + " cannot be longer than " + param + " characters"
					case "oneof":
						msg = fieldName + " must be one of the allowed values: " + param
					default:
						msg = fieldName + " failed validation on " + tag
					}

					errors[fieldName] = msg
				}
				c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}

			c.Abort()
			return
		}

		c.Set("validatedBody", obj)
		c.Next()
	}
}
