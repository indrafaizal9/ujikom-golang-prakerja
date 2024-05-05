package helpers

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func StructBinder(c *gin.Context, model interface{}) {
	contentType := c.Request.Header.Get("Content-Type")
	if contentType == "application/json" {
		c.ShouldBindJSON(&model)
	} else {
		c.ShouldBind(&model)
	}
}

func ValidateStruct(model interface{}) (bool, error) {
	_, err := govalidator.ValidateStruct(model)
	if err != nil {
		return false, err
	}
	return true, nil
}
