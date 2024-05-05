package services

import (
	"ujikom/database"
	"ujikom/pkg/helpers"
	"ujikom/pkg/models"
	"ujikom/pkg/resources"

	"github.com/gin-gonic/gin"
)

// Method List
// 1. Create Recipe
// 2. Update Recipe
// 3. Delete Recipe
// 4. Get Recipe
// 5. Get Recipes
// 6. Update Recipe Ingredient
// 7. Delete Recipe Ingredient
// 8. Add New Ingredient into Recipe

type RecipeService struct {
}

func (r *RecipeService) CreateRecipe(c *gin.Context, request models.RecipeCreate) {
	db := database.DB
	_ = db

	recipe := models.Recipe{
		Name:           request.Name,
		Description:    request.Description,
		PrepTime:       request.PrepTime,
		CookTime:       request.CookTime,
		AdditionalTime: request.AdditionalTime,
		TotalTime:      request.TotalTime,
		Servings:       request.Servings,
		Instructions:   request.Instructions,
		IsPublished:    request.IsPublished,
		Difficulty:     request.Difficulty,
		Tags:           request.Tags,
	}

	err := db.Create(&recipe).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	IngredientService := IngredientService{}
	Ingredient, err := IngredientService.IngredientsCreate(c, request.Ingredient, recipe.ID)
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	recipeResource := models.RecipeResource{}
	resources.RecipeMake(recipe, &recipeResource)
	recipeResource.Ingredients = Ingredient
	helpers.ResCreated(c, recipeResource)
}
