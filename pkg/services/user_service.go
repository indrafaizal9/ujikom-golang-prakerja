package services

import (
	"ujikom/database"
	"ujikom/pkg/helpers"
	"ujikom/pkg/models"
	"ujikom/pkg/resources"

	"github.com/gin-gonic/gin"
)

type UserService struct {
}

func (u *UserService) UpdateUser(c *gin.Context, parseID int, request models.UserUpdate) {
	db := database.DB

	var User models.User
	err := db.Where("id = ?", parseID).First(&User).Error
	if request.Username != "" {
		User.Username = request.Username
	}
	if request.Password != "" {
		User.Password = request.Password
	}
	if request.Role != "" {
		User.Role = request.Role
	}
	User.IsActive = request.IsActive

	if err != nil {
		helpers.ResNotFound(c, "User Not Found")
		return
	}

	err = db.Save(&User).Error
	if err != nil {
		helpers.ResInternalServerError(c, err.Error())
		return
	}

	userRersource := models.UserResource{}
	resources.UserMake(User, &userRersource)

	helpers.ResOK(c, userRersource)
}

func (u *UserService) DeleteUser(c *gin.Context, ID int) {
	db := database.DB

	User := models.User{}
	err := db.Where("id = ?", ID).First(&User).Error
	if err != nil {
		helpers.ResNotFound(c, err.Error())
		return
	}

	err = db.Delete(&User).Error
	if err != nil {
		helpers.ResInternalServerError(c, err.Error())
		return
	}

	helpers.ResOK(c, "User Deleted")
}

func (u *UserService) GetUser(c *gin.Context, ID int) {
	db := database.DB

	User := models.User{}
	helpers.StructBinder(c, &User)

	err := db.Where("id = ?", ID).First(&User).Error
	if err != nil {
		helpers.ResNotFound(c, err.Error())
		return
	}

	userRersource := models.UserResource{}
	resources.UserMake(User, &userRersource)

	helpers.ResOK(c, userRersource)
}

func (u *UserService) GetUsers(c *gin.Context) {
	db := database.DB

	var Users []models.User
	err := db.Find(&Users).Error
	if err != nil {
		helpers.ResInternalServerError(c, err.Error())
		return
	}

	if len(Users) == 0 {
		helpers.ResOK(c, nil)
		return
	}

	userResource := []models.UserResource{}
	resources.UserCollection(Users, &userResource)

	helpers.ResOK(c, userResource)
}
