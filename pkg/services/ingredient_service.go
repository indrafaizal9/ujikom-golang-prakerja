package services

import (
	"ujikom/database"
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

// func (i *IngredientService) IngredientCreate(c *gin.Context, request models.IngredientCreate) {
// 	Ingredient := models.Ingredient{
// 		Name: request.Name,
// 		Unit: request.Unit,
// 	}

// 	db := database.DB
// 	err := db.Create(&Ingredient).Error
// 	if err != nil {
// 		helpers.ResBadRequest(c, err.Error())
// 		return
// 	}
// }
