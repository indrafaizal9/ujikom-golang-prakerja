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
		UserID:         c.MustGet("user").(models.User).ID,
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

func (r *RecipeService) GetRecipes(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	db := database.DB

	recipes := []models.Recipe{}
	err := db.Where("user_id = ?", user.ID).Find(&recipes).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	recipeResources := []models.RecipeResource{}
	for _, recipe := range recipes {
		recipeResource := models.RecipeResource{}
		resources.RecipeMake(recipe, &recipeResource)
		recipeResources = append(recipeResources, recipeResource)
	}

	helpers.ResOK(c, recipeResources)
}

func (r *RecipeService) GetRecipe(c *gin.Context, id int) {
	db := database.DB

	recipe := models.Recipe{}
	err := db.Where("id = ?", id).First(&recipe).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	recipeResource := models.RecipeResource{}
	resources.RecipeMake(recipe, &recipeResource)

	IngredientService := IngredientService{}
	Ingredient, err := IngredientService.IngredientsGet(c, recipe.ID)
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	recipeResource.Ingredients = Ingredient
	helpers.ResOK(c, recipeResource)
}

func (r *RecipeService) UpdateRecipe(c *gin.Context, id int, request models.RecipeUpdate) {
	db := database.DB

	recipe := models.Recipe{}
	err := db.Where("id = ?", id).First(&recipe).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	if recipe.UserID != c.MustGet("user").(models.User).ID {
		helpers.ResBadRequest(c, "You are not authorized to update this recipe")
		return
	}

	if request.Name != "" {
		recipe.Name = request.Name
	}
	if request.Description != "" {
		recipe.Description = request.Description
	}
	if request.PrepTime != "" {
		recipe.PrepTime = request.PrepTime
	}
	if request.CookTime != "" {
		recipe.CookTime = request.CookTime
	}
	if request.AdditionalTime != "" {
		recipe.AdditionalTime = request.AdditionalTime
	}
	if request.TotalTime != "" {
		recipe.TotalTime = request.TotalTime
	}
	if request.Servings != "" {
		recipe.Servings = request.Servings
	}
	if request.Instructions != "" {
		recipe.Instructions = request.Instructions
	}

	if request.Difficulty != "" {
		recipe.Difficulty = request.Difficulty
	}
	if request.Tags != "" {
		recipe.Tags = request.Tags
	}
	if request.IsPublished != recipe.IsPublished {
		recipe.IsPublished = request.IsPublished
	}

	err = db.Save(&recipe).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	recipeResource := models.RecipeResource{}
	resources.RecipeMake(recipe, &recipeResource)
	helpers.ResOK(c, recipeResource)
}

func (r *RecipeService) DeleteRecipe(c *gin.Context, id int) {
	db := database.DB

	recipe := models.Recipe{}
	err := db.Where("id = ?", id).First(&recipe).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	if recipe.UserID != c.MustGet("user").(models.User).ID || c.MustGet("user").(models.User).Role != "admin"{
		helpers.ResBadRequest(c, "You are not authorized to delete this recipe")
		return
	}

	err = db.Delete(&recipe).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	helpers.ResOK(c, "Recipe deleted successfully")
}

func (r *RecipeService) GetPublicRecipes(c *gin.Context) {
	db := database.DB

	recipes := []models.Recipe{}
	err := db.Where("is_published = ?", true).Find(&recipes).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	recipeResources := []models.RecipeResource{}
	for _, recipe := range recipes {
		recipeResource := models.RecipeResource{}
		resources.RecipeMake(recipe, &recipeResource)
		recipeResources = append(recipeResources, recipeResource)
	}

	helpers.ResOK(c, recipeResources)
}

func (r *RecipeService) LikeRecipe(c *gin.Context, id int) {
	db := database.DB

	recipe := models.Recipe{}
	err := db.Where("id = ?", id).First(&recipe).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	recipe.LikeCount = recipe.LikeCount + 1
	err = db.Save(&recipe).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	helpers.ResOK(c, "Recipe liked successfully")
}