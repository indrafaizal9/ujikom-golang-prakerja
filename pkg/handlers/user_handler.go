package handlers

import (
	"strconv"
	"ujikom/pkg/helpers"
	"ujikom/pkg/models"
	"ujikom/pkg/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return UserHandler{userService: userService}
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

	u.userService.UpdateUser(c, parseID, request)
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	parseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	u.userService.DeleteUser(c, parseID)
}

func (u *UserHandler) GetUser(c *gin.Context) {
	parseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}
	u.userService.GetUser(c, parseID)
}

func (u *UserHandler) GetUsers(c *gin.Context) {
	u.userService.GetUsers(c)
}

func (u *UserHandler) GetRecipesByUser(c *gin.Context) {
	parseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}
	u.userService.GetRecipesByUser(c, parseID)
}

func (u *UserHandler) GetReviewsByUser(c *gin.Context) {
	parseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}
	u.userService.GetReviewsByUser(c, parseID)
}