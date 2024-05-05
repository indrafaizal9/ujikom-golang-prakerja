package handlers

import (
	"ujikom/pkg/helpers"
	"ujikom/pkg/models"
	"ujikom/pkg/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) AuthHandler {
	return AuthHandler{authService: authService}
}

func (a *AuthHandler) Login(c *gin.Context) {
	request := models.Login{}
	helpers.StructBinder(c, &request)
	_, errCreate := helpers.ValidateStruct(request)
	if errCreate != nil {
		helpers.ResBadRequest(c, errCreate.Error())
		return
	}
	a.authService.Login(c, request)
}

func (a *AuthHandler) Register(c *gin.Context) {
	request := models.UserCreate{}
	helpers.StructBinder(c, &request)

	_, errCreate := helpers.ValidateStruct(request)
	if errCreate != nil {
		helpers.ResBadRequest(c, errCreate.Error())
		return
	}

	a.authService.Register(c, request)
}

func (a *AuthHandler) Me(c *gin.Context) {
	a.authService.Me(c)
}
