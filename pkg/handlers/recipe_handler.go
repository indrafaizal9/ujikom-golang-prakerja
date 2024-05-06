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

func (r *RecipeHandler) CreateReview(c *gin.Context) {
	parseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	request := models.ReviewCreate{}
	helpers.StructBinder(c, &request)

	_, errCreate := helpers.ValidateStruct(request)
	if errCreate != nil {
		helpers.ResBadRequest(c, errCreate.Error())
		return
	}

	r.recipeService.CreateReview(c, parseID, request)
}

func (r *RecipeHandler) GetReviews(c *gin.Context) {
	parseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	r.recipeService.GetReviews(c, parseID)
}

func (r *RecipeHandler) UpdateReview(c *gin.Context) {
	parseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	parseReviewID, err := strconv.Atoi(c.Param("review_id"))
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	request := models.ReviewUpdate{}
	helpers.StructBinder(c, &request)

	_, errCreate := helpers.ValidateStruct(request)
	if errCreate != nil {
		helpers.ResBadRequest(c, errCreate.Error())
		return
	}

	r.recipeService.UpdateReview(c, parseID, parseReviewID, request)
}

func (r *RecipeHandler) DeleteReview(c *gin.Context) {
	parseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	parseReviewID, err := strconv.Atoi(c.Param("review_id"))
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	r.recipeService.DeleteReview(c, parseID, parseReviewID)
}

func (r *RecipeHandler) AddToCollection(c *gin.Context) {
	parseID, _ := strconv.Atoi(c.Param("id"))
	collectionID, err := strconv.Atoi(c.Param("collection_id"))
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	r.recipeService.AddToCollection(c, parseID, collectionID)
}

func (r *RecipeHandler) CreateTag(c *gin.Context) {
	request := models.TagsAndLabels{}
	helpers.StructBinder(c, &request)

	_, errCreate := helpers.ValidateStruct(request)
	if errCreate != nil {
		helpers.ResBadRequest(c, errCreate.Error())
		return
	}

	r.recipeService.CreateTag(c, request)
}

func (r *RecipeHandler) GetTags(c *gin.Context) {
	r.recipeService.GetTags(c)
}

func (r *RecipeHandler) DeleteTag(c *gin.Context) {
	parseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	r.recipeService.DeleteTag(c, parseID)
}

func (r *RecipeHandler) CreateLabel(c *gin.Context) {
	request := models.TagsAndLabels{}
	helpers.StructBinder(c, &request)

	_, errCreate := helpers.ValidateStruct(request)
	if errCreate != nil {
		helpers.ResBadRequest(c, errCreate.Error())
		return
	}

	r.recipeService.CreateLabel(c, request)
}

func (r *RecipeHandler) GetLabels(c *gin.Context) {
	r.recipeService.GetLabels(c)
}

func (r *RecipeHandler) DeleteLabel(c *gin.Context) {
	parseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	r.recipeService.DeleteLabel(c, parseID)
}

func (r *RecipeHandler) SearchRecipe(c *gin.Context) {
	searchModel := models.SearchRecipe{}
	helpers.StructBinder(c, &searchModel)

	_, errCreate := helpers.ValidateStruct(searchModel)
	if errCreate != nil {
		helpers.ResBadRequest(c, errCreate.Error())
		return
	}

	r.recipeService.SearchRecipes(c, searchModel)
}
