package middlewares

import (
	"ujikom/database"
	"ujikom/pkg/helpers"
	"ujikom/pkg/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			helpers.ResUnauthorized(c, err.Error())
			c.Abort()
			return
		}

		c.Set("userdata", verifyToken)
		User(c)
		c.Next()
	}
}

func User(c *gin.Context) models.User {
	userData := c.MustGet("userdata").(jwt.MapClaims)
	userID := int(userData["id"].(float64))
	user := models.User{}
	database.DB.Where("id = ?", userID).First(&user)
	c.Set("user", user)
	return user
}
