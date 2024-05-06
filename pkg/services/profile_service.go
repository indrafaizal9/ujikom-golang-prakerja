package services

import (
	"ujikom/database"
	"ujikom/pkg/helpers"
	"ujikom/pkg/models"
	"ujikom/pkg/resources"

	"github.com/gin-gonic/gin"
)

type ProfileService struct {
}

func (s *ProfileService) GetProfileByUserID(c *gin.Context, userID uint) {
	db := database.DB

	profile := models.Profile{}
	err := db.Where("user_id = ?", userID).First(&profile).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	//get username
	user := models.User{}
	err = db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	profileResource := models.ProfileResource{}
	resources.ProfileMake(profile, &profileResource)
	profileResource.Usermame = user.Username
	helpers.ResOK(c, profileResource)
}

func (s *ProfileService) CreateProfile(c *gin.Context, userID uint) {
	db := database.DB

	profile := models.Profile{}
	err := db.Where("user_id = ?", userID).First(&profile).Error
	if err == nil {
		helpers.ResBadRequest(c, "Profile already exists")
		return
	}

	request := models.ProfileCreate{}
	err = c.Bind(&request)
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	profile = models.Profile{
		UserID:   userID,
		FullName: request.FullName,
		Email:    request.Email,
		Gender:  request.Gender,
		Address:  request.Address,
		Photo:    request.Photo,
	}

	err = db.Create(&profile).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	user := models.User{}
	err = db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	profileResource := models.ProfileResource{}
	resources.ProfileMake(profile, &profileResource)
	profileResource.Usermame = user.Username
	helpers.ResOK(c, profileResource)
}

func (s *ProfileService) UpdateProfile(c *gin.Context, userID uint) {
	db := database.DB

	profile := models.Profile{}
	err := db.Where("user_id = ?", userID).First(&profile).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	if profile.UserID != userID {
		helpers.ResBadRequest(c, "You are not authorized to update this profile")
		return
	}

	request := models.ProfileUpdate{}
	err = c.Bind(&request)
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	if request.FullName != "" {
		profile.FullName = request.FullName
	}
	if request.Email != "" {
		profile.Email = request.Email
	}
	if request.Gender != "" {
		profile.Gender = request.Gender
	}
	if request.Address != "" {
		profile.Address = request.Address
	}
	if request.Photo != "" {
		profile.Photo = request.Photo
	}

	err = db.Save(&profile).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	user := models.User{}
	err = db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	profileResource := models.ProfileResource{}
	resources.ProfileMake(profile, &profileResource)
	profileResource.Usermame = user.Username
	helpers.ResOK(c, profileResource)
}

func (s *ProfileService) GetRecipesByUserID(c *gin.Context, userID uint) {
	db := database.DB

	recipes := []models.Recipe{}
	err := db.Where("user_id = ?", userID).Find(&recipes).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	recipeResources := []models.RecipeResource{}
	resources.RecipeCollection(recipes, &recipeResources)
	helpers.ResOK(c, recipeResources)
}

func (s *ProfileService) GetCollections(c *gin.Context, userID uint) {
	db := database.DB

	collections := []models.Collection{}
	err := db.Where("user_id = ?", userID).Find(&collections).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	collectionResources := []models.CollectionResource{}
	resources.CollectionCollection(collections, &collectionResources)
	helpers.ResOK(c, collectionResources)
}

func (s *ProfileService) GetCollection(c *gin.Context, userID uint, collectionID uint) {
	db := database.DB

	collection := models.Collection{}
	err := db.Where("user_id = ? AND id = ?", userID, collectionID).First(&collection).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	recipes := []models.Recipe{}
	err = db.Table("recipes").Select("recipes.*").Joins("JOIN collection_recipes_pivot ON collection_recipes_pivot.recipe_id = recipes.id").Where("collection_recipes_pivot.collection_id = ?", collectionID).Find(&recipes).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}
	recipeResource := []models.RecipeResource{}
	resources.RecipeCollection(recipes, &recipeResource)

	collectionResource := models.CollectionResource{}
	resources.CollectionMake(collection, &collectionResource)
	collectionResource.RecipeCount = uint(len(recipes))
	collectionResource.Recipes = recipeResource
	helpers.ResOK(c, collectionResource)
}

func (s *ProfileService) CreateCollection(c *gin.Context, userID uint, request models.CollectionCreate) {
	db := database.DB

	collection := models.Collection{
		UserID:      userID,
		Name:        request.Name,
		Description: request.Description,
	}

	err := db.Create(&collection).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	collectionResource := models.CollectionResource{}
	resources.CollectionMake(collection, &collectionResource)
	helpers.ResOK(c, collectionResource)
}

func (s *ProfileService) UpdateCollection(c *gin.Context, userID uint, collectionID uint, request models.CollectionUpdate) {
	db := database.DB

	collection := models.Collection{}
	err := db.Where("user_id = ? AND id = ?", userID, collectionID).First(&collection).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	if collection.UserID != userID {
		helpers.ResBadRequest(c, "You are not authorized to update this collection")
		return
	}

	if request.Name != "" {
		collection.Name = request.Name
	}
	if request.Public {
		collection.Public = request.Public
	}
	if request.Description != "" {
		collection.Description = request.Description
	}

	err = db.Save(&collection).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	collectionResource := models.CollectionResource{}
	resources.CollectionMake(collection, &collectionResource)
	helpers.ResOK(c, collectionResource)
}

func (s *ProfileService) DeleteCollection(c *gin.Context, userID uint, collectionID uint) {
	db := database.DB

	collection := models.Collection{}
	err := db.Where("user_id = ? AND id = ?", userID, collectionID).First(&collection).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	if collection.UserID != userID {
		helpers.ResBadRequest(c, "You are not authorized to delete this collection")
		return
	}

	err = db.Where("id = ?", collectionID).Delete(&collection).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	helpers.ResOK(c, "Collection deleted")
}