package handlers

import (
	"strconv"
	"ujikom/pkg/helpers"
	"ujikom/pkg/models"
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
	parseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}
	request := models.UserUpdate{}
	helpers.StructBinder(c, &request)

	_, errCreate := helpers.ValidateStruct(request)
	if errCreate != nil {
		helpers.ResBadRequest(c, errCreate.Error())
		return
	}

	u.userHandler.UpdateUser(c, parseID, request)
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	parseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	u.userHandler.DeleteUser(c, parseID)
}

func (u *UserHandler) GetUser(c *gin.Context) {
	parseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}
	u.userHandler.GetUser(c, parseID)
}

func (u *UserHandler) GetUsers(c *gin.Context) {
	u.userHandler.GetUsers(c)
}
