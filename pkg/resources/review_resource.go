package resources

import "ujikom/pkg/models"

func ReviewMake(source models.Review, target *models.ReviewResource) {
	target.ID = source.ID
	target.UserID = source.UserID
	target.RecipeID = source.RecipeID
	target.Rating = source.Rating
	target.Review = source.Review
}

func ReviewCollection(source []models.Review, target *[]models.ReviewResource) {
	for _, item := range source {
		var reviewResource models.ReviewResource
		ReviewMake(item, &reviewResource)
		*target = append(*target, reviewResource)
	}
}
