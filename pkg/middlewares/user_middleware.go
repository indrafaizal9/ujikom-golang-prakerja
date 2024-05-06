package middlewares

import (
	"ujikom/pkg/helpers"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RoleUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userData := c.MustGet("userdata").(jwt.MapClaims)
		role := userData["role"].(string)

		if role != "user" {
			helpers.ResUnauthorized(c, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func AllowedRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userData := c.MustGet("userdata").(jwt.MapClaims)
		userRole := userData["role"].(string)

		if userRole != role {
			helpers.ResUnauthorized(c, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
