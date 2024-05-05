package handlers

import (
	"fmt"
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
