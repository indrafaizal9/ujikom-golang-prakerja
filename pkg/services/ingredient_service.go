package services

import (
	"ujikom/database"
	"ujikom/pkg/helpers"
	"ujikom/pkg/models"
	"ujikom/pkg/resources"

	"github.com/gin-gonic/gin"
)

type IngredientService struct {
}

func (i *IngredientService) IngredientsCreate(c *gin.Context, request []models.IngredientCreate, recipeID uint) ([]models.IngredientResource, error) {
	var Ingredients []models.Ingredient
	for _, item := range request {
		Ingredient := models.Ingredient{
			Name:        item.Name,
			Description: item.Description,
			RecipeID:    recipeID,
			Amount:      item.Amount,
			Unit:        item.Unit,
		}
		Ingredients = append(Ingredients, Ingredient)
	}

	db := database.DB
	err := db.Create(&Ingredients).Error

	var IngredientsResource []models.IngredientResource
	resources.IngredientCollection(Ingredients, &IngredientsResource)
	return IngredientsResource, err
}

func (i *IngredientService) IngredientsGet(c *gin.Context, recipeID uint) ([]models.IngredientResource, error) {
	db := database.DB
	var Ingredients []models.Ingredient
	err := db.Where("recipe_id = ?", recipeID).Find(&Ingredients).Error

	var IngredientsResource []models.IngredientResource
	resources.IngredientCollection(Ingredients, &IngredientsResource)
	return IngredientsResource, err
}

func (i *IngredientService) AddIngredient(c *gin.Context, request models.IngredientCreate, recipeID int) {
	db := database.DB

	Ingredient := models.Ingredient{
		Name:        request.Name,
		Description: request.Description,
		RecipeID:    uint(recipeID),
		Amount:      request.Amount,
		Unit:        request.Unit,
	}

	err := db.Create(&Ingredient).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	IngredientResource := models.IngredientResource{}
	resources.IngredientMake(Ingredient, &IngredientResource)
	helpers.ResCreated(c, IngredientResource)
}

func (i *IngredientService) UpdateIngredient(c *gin.Context, request models.IngredientUpdate, recipeID, ingredientID int) {
	db := database.DB
	userData := c.MustGet("user").(models.User)
	Ingredient := models.Ingredient{}
	err := db.Where("id = ? AND recipe_id = ?", ingredientID, recipeID).First(&Ingredient).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	//check if the user is the owner of the recipe
	recipe := models.Recipe{}
	err = db.Where("id = ?", recipeID).First(&recipe).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	if recipe.UserID != userData.ID {
		helpers.ResBadRequest(c, "You are not the owner of this recipe")
		return
	}

	if request.Name != "" {
		Ingredient.Name = request.Name
	}
	if request.Description != "" {
		Ingredient.Description = request.Description
	}
	if request.Amount != "" {
		Ingredient.Amount = request.Amount
	}
	if request.Unit != "" {
		Ingredient.Unit = request.Unit
	}

	err = db.Save(&Ingredient).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	IngredientResource := models.IngredientResource{}
	resources.IngredientMake(Ingredient, &IngredientResource)
	helpers.ResOK(c, IngredientResource)
}

func (i *IngredientService) DeleteIngredient(c *gin.Context, recipeID, ingredientID int) {

	db := database.DB
	userData := c.MustGet("user").(models.User)
	Ingredient := models.Ingredient{}
	err := db.Where("id = ? AND recipe_id = ?", ingredientID, recipeID).First(&Ingredient).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	//check if the user is the owner of the recipe
	recipe := models.Recipe{}
	err = db.Where("id = ?", recipeID).First(&recipe).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	if recipe.UserID != userData.ID || userData.Role != "admin" {
		helpers.ResBadRequest(c, "You are not the owner of this recipe")
		return
	}

	err = db.Delete(&Ingredient).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	helpers.ResNoContent(c)
}
