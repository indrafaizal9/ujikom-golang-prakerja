package services

import (
	"ujikom/database"
	"ujikom/pkg/helpers"
	"ujikom/pkg/models"
	"ujikom/pkg/resources"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
}

func (a *AuthService) Login(c *gin.Context, request models.Login) {
	db := database.DB
	_ = db

	userData := models.User{
		Username: request.Username,
		Password: request.Password,
	}

	var User models.User
	err := db.Where("username = ? AND is_active IS TRUE", userData.Username).First(&User).Error
	if err != nil {
		helpers.ResUnauthorized(c, "Invalid Username or Password")
		return
	}

	comparePassword := helpers.ComparePassword(userData.Password, request.Password)
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

func (a *AuthService) Register(c *gin.Context, request models.UserCreate) {
	db := database.DB
	_ = db

	var user models.User
	userExist := db.Where("username = ?", request.Username).First(&user).Error
	if userExist == nil {
		helpers.ResBadRequest(c, "Username already exist")
		return
	}

	userData := models.User{
		Username: request.Username,
		Password: helpers.HashPassword(request.Password),
		Role:     request.Role,
		IsActive: true,
	}

	err := db.Create(&userData).Error
	if err != nil {
		helpers.ResInternalServerError(c, err.Error())
		return
	}

	userResource := models.UserResource{}
	resources.UserMake(userData, &userResource)
	helpers.ResCreated(c, userResource)
}
