package resources

import "ujikom/pkg/models"

func ProfileMake(profile models.Profile, profileResource *models.ProfileResource) {
	profileResource.ID = profile.ID
	profileResource.UserID = profile.UserID
	profileResource.FullName = profile.FullName
	profileResource.Email = profile.Email
	profileResource.Gender = profile.Gender
	profileResource.Address = profile.Address
	profileResource.Photo = profile.Photo
}

func ProfileCollection(profiles []models.Profile, profileResources *[]models.ProfileResource) {
	for _, profile := range profiles {
		profileResource := models.ProfileResource{}
		ProfileMake(profile, &profileResource)
		*profileResources = append(*profileResources, profileResource)
	}
}
