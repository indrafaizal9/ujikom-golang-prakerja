package middlewares

import (
	"ujikom/pkg/helpers"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc{
	return func(c *gin.Context){
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			helpers.ResUnauthorized(c, err.Error())
			c.Abort()
			return
		}
		c.Set("userdata", verifyToken)
		c.Next()
	}
}