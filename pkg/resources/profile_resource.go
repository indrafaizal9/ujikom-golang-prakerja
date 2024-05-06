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