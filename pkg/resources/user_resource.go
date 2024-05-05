package resources

import "ujikom/pkg/models"

func UserMake(source models.User, target *models.UserResource) {
	target.ID = source.ID
	target.Username = source.Username
	target.Role = source.Role
	target.IsActive = source.IsActive
	target.CreatedAt = source.CreatedAt.Format("2006-01-02 15:04:05")
}

func UserCollection(source []models.User, target *[]models.UserResource) {
	for _, item := range source {
		var userResource models.UserResource
		UserMake(item, &userResource)
		*target = append(*target, userResource)
	}
}