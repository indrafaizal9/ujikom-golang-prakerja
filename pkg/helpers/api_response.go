package helpers

import (
	"github.com/gin-gonic/gin"
)

type ResponseSuccess struct {
	Status  int         `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseError struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
}

func ResOK(c *gin.Context, data interface{}) {
	c.JSON(200, ResponseSuccess{Status: 200, Data: data})
}
func ResError(c *gin.Context, message string) {
	c.JSON(400, ResponseError{Status: 400, Message: message})
}
func ResCreated(c *gin.Context, data interface{}) {
	c.JSON(201, ResponseSuccess{Status: 201, Data: data})
}
func ResBadRequest(c *gin.Context, message string) {
	c.JSON(400, ResponseError{Status: 400, Message: message})
}
func ResUnauthorized(c *gin.Context, message string) {
	c.JSON(401, ResponseError{Status: 401, Message: message})
}
func ResForbidden(c *gin.Context, message string) {
	c.JSON(403, ResponseError{Status: 403, Message: message})
}
func ResNotFound(c *gin.Context, message string) {
	c.JSON(404, ResponseError{Status: 404, Message: message})
}
func ResInternalServerError(c *gin.Context, message string) {
	c.JSON(500, ResponseError{Status: 500, Message: message})
}
func ResServiceUnavailable(c *gin.Context, message string) {
	c.JSON(503, ResponseError{Status: 503, Message: message})
}
func ResConflict(c *gin.Context, message string) {
	c.JSON(409, ResponseError{Status: 409, Message: message})
}
func ResNoContent(c *gin.Context) {
	c.JSON(204, nil)
}

func ResCustom(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}
