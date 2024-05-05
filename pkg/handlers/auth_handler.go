package handlers

import (
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
	a.authService.Login(c)
}

func (a *AuthHandler) Register(c *gin.Context) {
	a.authService.Register(c)
}
