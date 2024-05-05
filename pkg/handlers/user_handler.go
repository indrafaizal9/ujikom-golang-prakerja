package handlers

import (
	"ujikom/pkg/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userHandler services.UserService
}

func NewUserHandler(userHandler services.UserService) UserHandler {
	return UserHandler{userHandler: userHandler}
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	u.userHandler.UpdateUser(c)
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	u.userHandler.DeleteUser(c)
}

func (u *UserHandler) GetUser(c *gin.Context) {
	u.userHandler.GetUser(c)
}

func (u *UserHandler) GetUsers(c *gin.Context) {
	u.userHandler.GetUsers(c)
}