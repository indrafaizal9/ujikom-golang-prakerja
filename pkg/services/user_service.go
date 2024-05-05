package services

import (
	"ujikom/database"
	"ujikom/pkg/helpers"
	"ujikom/pkg/models"

	"github.com/gin-gonic/gin"
)

type UserService struct {
}

func (u *UserService) UpdateUser(c *gin.Context) {
	db := database.DB
	_ = db

	User := models.UserUpdate{}
	helpers.StructBinder(c, &User)

	err := db.Save(&User).Error
	if err != nil {
		helpers.ResInternalServerError(c, err.Error())
		return
	}

	helpers.ResOK(c, User)
}

func (u *UserService) DeleteUser(c *gin.Context) {
	db := database.DB
	_ = db

	User := models.User{}
	helpers.StructBinder(c, &User)

	err := db.Delete(&User).Error
	if err != nil {
		helpers.ResInternalServerError(c, err.Error())
		return
	}

	helpers.ResOK(c, User)
}

func (u *UserService) GetUser(c *gin.Context) {
	db := database.DB
	_ = db

	User := models.User{}
	helpers.StructBinder(c, &User)

	err := db.First(&User).Error
	if err != nil {
		helpers.ResNotFound(c, err.Error())
		return
	}

	helpers.ResOK(c, User)
}

func (u *UserService) GetUsers(c *gin.Context) {
	db := database.DB
	_ = db

	var Users []models.User
	err := db.Find(&Users).Error
	if err != nil {
		helpers.ResInternalServerError(c, err.Error())
		return
	}

	helpers.ResOK(c, Users)
}
