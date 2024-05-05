package services

import (
	"ujikom/database"
	"ujikom/pkg/helpers"
	"ujikom/pkg/models"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
}

func (a *AuthService) Login(c *gin.Context) {
	db := database.DB
	_ = db

	User := models.User{}
	helpers.StructBinder(c, &User)

	password := User.Password

	err := db.Where("username = ? AND is_active IS TRUE", User.Username).First(&User).Error
	if err != nil {
		helpers.ResUnauthorized(c, "Invalid Username or Password")
		return
	}

	comparePassword := helpers.ComparePassword(User.Password, password)
	if !comparePassword {
		helpers.ResUnauthorized(c, "Invalid Username or Password")
		return
	}

	token, err := helpers.GenerateToken(User.ID, User.Username)
	if err != nil {
		helpers.ResInternalServerError(c, err.Error())
		return
	}

	helpers.ResOK(c, map[string]interface{}{
		"token": token,
	})
}

func (a *AuthService) Register(c *gin.Context) {
	db := database.DB
	_ = db

	User := models.UserCreate{}
	helpers.StructBinder(c, &User)

	_, errCreate := helpers.ValidateStruct(User)
	if errCreate != nil {
		helpers.ResBadRequest(c, errCreate.Error())
		return
	}

	var user models.User
	userExist := db.Where("username = ?", User.Username).First(&user).Error
	if userExist == nil {
		helpers.ResBadRequest(c, "Username already exist")
		return
	}

	userData := models.User{
		Username: User.Username,
		Password: helpers.HashPassword(User.Password),
		Role:     User.Role,
		IsActive: true,
	}

	err := db.Create(&userData).Error
	if err != nil {
		helpers.ResInternalServerError(c, err.Error())
		return
	}

	helpers.ResCreated(c, userData)
}
