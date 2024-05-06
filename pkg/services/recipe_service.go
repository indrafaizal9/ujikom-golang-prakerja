package services

import (
	"ujikom/database"
	"ujikom/pkg/helpers"
	"ujikom/pkg/models"
	"ujikom/pkg/resources"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

	if recipe.UserID != c.MustGet("user").(models.User).ID || c.MustGet("user").(models.User).Role != "admin" {
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

	// Check if user already liked the recipe
	user := c.MustGet("user").(models.User)
	like := models.Likes{}
	err = db.Where("user_id = ? AND recipe_id = ?", user.ID, recipe.ID).First(&like).Error

	// Start a transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err == nil { // If user already liked the recipe, unlike it
		if err := tx.Delete(&like).Error; err != nil {
			tx.Rollback()
			helpers.ResBadRequest(c, err.Error())
			return
		}
		// Decrement like count in the database
		if err := tx.Model(&recipe).UpdateColumn("like_count", gorm.Expr("like_count - 1")).Error; err != nil {
			tx.Rollback()
			helpers.ResBadRequest(c, err.Error())
			return
		}
		helpers.ResOK(c, "Recipe unliked")
	} else { // If user hasn't liked the recipe yet, like it
		like = models.Likes{
			UserID:   user.ID,
			RecipeID: recipe.ID,
		}
		if err := tx.Create(&like).Error; err != nil {
			tx.Rollback()
			helpers.ResBadRequest(c, err.Error())
			return
		}
		// Increment like count in the database
		if err := tx.Model(&recipe).UpdateColumn("like_count", gorm.Expr("like_count + 1")).Error; err != nil {
			tx.Rollback()
			helpers.ResBadRequest(c, err.Error())
			return
		}
		helpers.ResOK(c, "Recipe liked")
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}
}

func (r *RecipeService) CreateReview(c *gin.Context, id int, request models.ReviewCreate) {
	db := database.DB

	recipe := models.Recipe{}
	err := db.Where("id = ?", id).First(&recipe).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	// Check if user already reviewed the recipe
	if recipe.UserID == c.MustGet("user").(models.User).ID {
		helpers.ResBadRequest(c, "You can't review your own recipe")
		return
	}

	review := models.Review{}
	err = db.Where("user_id = ? AND recipe_id = ?", c.MustGet("user").(models.User).ID, recipe.ID).First(&review).Error
	if err == nil {
		helpers.ResBadRequest(c, "You already reviewed this recipe")
		return
	}

	user := c.MustGet("user").(models.User)
	reviewData := models.Review{
		UserID:   user.ID,
		RecipeID: recipe.ID,
		Rating:   request.Rating,
		Review:   request.Review,
	}

	err = db.Create(&reviewData).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	//calculate average rating
	var totalRating float64
	var reviewCount int64

	db.Model(&models.Review{}).Where("recipe_id = ?", recipe.ID).Count(&reviewCount)
	db.Model(&models.Review{}).Where("recipe_id = ?", recipe.ID).Select("SUM(rating)").Row().Scan(&totalRating)
	averageRating := totalRating / float64(reviewCount)

	// Update recipe rating
	err = db.Model(&recipe).Update("rating", averageRating).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	reviewResource := models.ReviewResource{}
	resources.ReviewMake(reviewData, &reviewResource)

	helpers.ResCreated(c, reviewResource)
}

func (r *RecipeService) UpdateReview(c *gin.Context, recipeID int, reviewID int, request models.ReviewUpdate) {
	db := database.DB

	review := models.Review{}
	err := db.Where("id = ? and recipe_id = ?", reviewID, recipeID).First(&review).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	if review.UserID != c.MustGet("user").(models.User).ID {
		helpers.ResBadRequest(c, "You are not authorized to update this review")
		return
	}

	updateRating := false

	if request.Rating != 0 {
		review.Rating = request.Rating
		updateRating = true
	}
	if request.Review != "" {
		review.Review = request.Review
	}

	err = db.Save(&review).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	if updateRating {
		//calculate average rating
		var totalRating float64
		var reviewCount int64

		db.Model(&models.Review{}).Where("recipe_id = ?", review.RecipeID).Count(&reviewCount)
		db.Model(&models.Review{}).Where("recipe_id = ?", review.RecipeID).Select("SUM(rating)").Row().Scan(&totalRating)
		averageRating := totalRating / float64(reviewCount)

		// Update recipe rating
		err = db.Model(&models.Recipe{}).Where("id = ?", review.RecipeID).Update("rating", averageRating).Error
		if err != nil {
			helpers.ResBadRequest(c, err.Error())
			return
		}
	}

	reviewResource := models.ReviewResource{}
	resources.ReviewMake(review, &reviewResource)

	helpers.ResOK(c, reviewResource)
}

func (r *RecipeService) GetReviews(c *gin.Context, id int) {
	db := database.DB

	reviews := []models.Review{}
	err := db.Where("recipe_id = ?", id).Find(&reviews).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	reviewResources := []models.ReviewResource{}
	resources.ReviewCollection(reviews, &reviewResources)

	helpers.ResOK(c, reviewResources)
}

func (r *RecipeService) DeleteReview(c *gin.Context, recipeID int, reviewID int) {
	db := database.DB

	review := models.Review{}
	err := db.Where("id = ? and recipe_id = ?", reviewID, recipeID).First(&review).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	if review.UserID != c.MustGet("user").(models.User).ID {
		helpers.ResBadRequest(c, "You are not authorized to delete this review")
		return
	}

	//calculate average rating
	var totalRating float64
	var reviewCount int64

	db.Model(&models.Review{}).Where("recipe_id = ?", review.RecipeID).Count(&reviewCount)
	db.Model(&models.Review{}).Where("recipe_id = ?", review.RecipeID).Select("SUM(rating)").Row().Scan(&totalRating)
	averageRating := totalRating / float64(reviewCount)

	// Update recipe rating
	err = db.Model(&models.Recipe{}).Where("id = ?", review.RecipeID).Update("rating", averageRating).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	err = db.Delete(&review).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	helpers.ResOK(c, "Review deleted successfully")
}

func (r *RecipeService) AddToCollection(c *gin.Context, recipeID int, collectionID int) {
	db := database.DB

	recipe := models.Recipe{}
	err := db.Where("id = ?", recipeID).First(&recipe).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	collection := models.Collection{}
	err = db.Where("id = ?", collectionID).First(&collection).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	recipeCollection := models.CollectionRecipesPivot{}
	err = db.Where("recipe_id = ? AND collection_id = ?", recipe.ID, collection.ID).First(&recipeCollection).Error
	if err == nil {
		helpers.ResBadRequest(c, "Recipe already in collection")
		return
	}

	recipeCollection = models.CollectionRecipesPivot{
		RecipeID:     uint(recipeID),
		CollectionID: uint(collection.ID),
	}

	err = db.Create(&recipeCollection).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	helpers.ResCreated(c, "Recipe added to collection")
}

func (r *RecipeService) CreateTag(c *gin.Context, request models.TagsAndLabels) {
	db := database.DB

	tag := models.TagsAndLabels{
		Name: request.Name,
		Type: "tag",
	}

	err := db.Create(&tag).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	helpers.ResCreated(c, tag)
}

func (r *RecipeService) GetTags(c *gin.Context) {
	db := database.DB

	tags := []models.TagsAndLabels{}
	err := db.Where("type = ?", "tag").Find(&tags).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	helpers.ResOK(c, tags)
}

func (r *RecipeService) DeleteTag(c *gin.Context, id int) {
	db := database.DB

	tag := models.TagsAndLabels{}
	err := db.Where("id = ?", id).First(&tag).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	err = db.Delete(&tag).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	helpers.ResOK(c, "Tag deleted successfully")
}

func (r *RecipeService) CreateLabel(c *gin.Context, request models.TagsAndLabels) {
	db := database.DB

	label := models.TagsAndLabels{
		Name: request.Name,
		Type: "label",
	}

	err := db.Create(&label).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	helpers.ResCreated(c, label)
}

func (r *RecipeService) GetLabels(c *gin.Context) {
	db := database.DB

	labels := []models.TagsAndLabels{}
	err := db.Where("type = ?", "label").Find(&labels).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	helpers.ResOK(c, labels)
}

func (r *RecipeService) DeleteLabel(c *gin.Context, id int) {
	db := database.DB

	label := models.TagsAndLabels{}
	err := db.Where("id = ?", id).First(&label).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	err = db.Delete(&label).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	helpers.ResOK(c, "Label deleted successfully")
}

func (r *RecipeService) SearchRecipes(c *gin.Context, request models.SearchRecipe) {
	db := database.DB

	recipes := []models.Recipe{}
	err := db.Where("name LIKE ? OR tag LIKE", "%"+request.Name+"%", "%"+request.Tags+"%").Find(&recipes).Error
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}

	recipeResources := []models.RecipeResource{}
	resources.RecipeCollection(recipes, &recipeResources)
	helpers.ResOK(c, recipeResources)
}