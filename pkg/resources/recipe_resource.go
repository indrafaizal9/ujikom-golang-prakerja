package resources

import "ujikom/pkg/models"

func RecipeMake(source models.Recipe, target *models.RecipeResource) {
	target.ID = source.ID
	target.Name = source.Name
	target.Description = source.Description
	target.PrepTime = source.PrepTime + " minutes"
	target.CookTime = source.CookTime + " minutes"
	target.TotalTime = source.TotalTime + " minutes"
	target.Servings = source.Servings
	target.LikeCount = source.LikeCount
	target.Instructions = source.Instructions
	target.Rating = source.Rating
	target.IsPublished = source.IsPublished
	target.Difficulty = source.Difficulty
	target.Tags = source.Tags
	target.Ingredients = nil
}

func RecipeCollection(source []models.Recipe, target *[]models.RecipeResource) {
	for _, item := range source {
		var recipeResource models.RecipeResource
		RecipeMake(item, &recipeResource)
		*target = append(*target, recipeResource)
	}
}
