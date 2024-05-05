package handlers

import (
	"fmt"
	"strconv"
	"ujikom/pkg/helpers"
	"ujikom/pkg/models"
	"ujikom/pkg/services"

	"github.com/gin-gonic/gin"
)

type RecipeHandler struct {
	recipeService services.RecipeService
}

func NewRecipeHandler(recipeService services.RecipeService) RecipeHandler {
	return RecipeHandler{recipeService: recipeService}
}

func (r *RecipeHandler) CreateRecipe(c *gin.Context) {
	request := models.RecipeCreate{}
	helpers.StructBinder(c, &request)

	fmt.Println("Request:", request)

	_, errCreate := helpers.ValidateStruct(request)
	if errCreate != nil {
		fmt.Println("Error:", errCreate)
		helpers.ResBadRequest(c, errCreate.Error())
		return
	}
	fmt.Println("Request:", request)

	r.recipeService.CreateRecipe(c, request)
}

func (r *RecipeHandler) GetRecipes(c *gin.Context) {
	r.recipeService.GetRecipes(c)
}

func (r *RecipeHandler) GetRecipe(c *gin.Context) {
	parseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}
	r.recipeService.GetRecipe(c, parseID)
}

func (r *RecipeHandler) UpdateRecipe(c *gin.Context) {
	parseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	request := models.RecipeUpdate{}
	helpers.StructBinder(c, &request)

	_, errCreate := helpers.ValidateStruct(request)
	if errCreate != nil {
		helpers.ResBadRequest(c, errCreate.Error())
		return
	}

	r.recipeService.UpdateRecipe(c, parseID, request)
}

func (r *RecipeHandler) DeleteRecipe(c *gin.Context) {
	parseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	r.recipeService.DeleteRecipe(c, parseID)
}

func (r *RecipeHandler) GetPublicRecipes(c *gin.Context) {
	r.recipeService.GetPublicRecipes(c)
}

func (r *RecipeHandler) LikeRecipe(c *gin.Context) {
	parseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	r.recipeService.LikeRecipe(c, parseID)
}