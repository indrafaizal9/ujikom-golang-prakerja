package handlers

import (
	"strconv"
	"ujikom/pkg/helpers"
	"ujikom/pkg/models"
	"ujikom/pkg/services"

	"github.com/gin-gonic/gin"
)

type IngredientHandler struct {
	ingredientService services.IngredientService
}

func NewIngredientHandler(ingredientService services.IngredientService) IngredientHandler {
	return IngredientHandler{ingredientService: ingredientService}
}

func (i *IngredientHandler) AddIngredient(c *gin.Context) {
	request := models.IngredientCreate{}
	helpers.StructBinder(c, &request)

	_, errCreate := helpers.ValidateStruct(request)
	if errCreate != nil {
		helpers.ResBadRequest(c, errCreate.Error())
		return
	}

	recipeID, _ := strconv.Atoi(c.Param("id"))
	i.ingredientService.AddIngredient(c, request, recipeID)
}

func (i *IngredientHandler) UpdateIngredient(c *gin.Context) {
	recipeID, _ := strconv.Atoi(c.Param("id"))
	ingredientID, _ := strconv.Atoi(c.Param("ingredient_id"))

	request := models.IngredientUpdate{}
	helpers.StructBinder(c, &request)

	_, errCreate := helpers.ValidateStruct(request)
	if errCreate != nil {
		helpers.ResBadRequest(c, errCreate.Error())
		return
	}

	i.ingredientService.UpdateIngredient(c, request, recipeID, ingredientID)
}

func (i *IngredientHandler) DeleteIngredient(c *gin.Context) {
	recipeID, _ := strconv.Atoi(c.Param("id"))
	ingredientID, _ := strconv.Atoi(c.Param("ingredient_id"))

	i.ingredientService.DeleteIngredient(c, recipeID, ingredientID)
}
