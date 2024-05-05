package resources

import "ujikom/pkg/models"

func IngredientMake(source models.Ingredient, target *models.IngredientResource) {
	target.ID = source.ID
	target.Name = source.Name
	target.Description = source.Description
	target.Amount = source.Amount
	target.Unit = source.Unit
}

func IngredientCollection(sources []models.Ingredient, targets *[]models.IngredientResource) {
	for _, ingredient := range sources {
		target := models.IngredientResource{}
		IngredientMake(ingredient, &target)
		*targets = append(*targets, target)
	}
}
