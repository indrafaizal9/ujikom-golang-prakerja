package handlers

import (
	"ujikom/pkg/services"
)

type ProfileHandler struct {
	profileService services.ProfileService
}

func NewProfileHandler(profileService services.ProfileService) ProfileHandler {
	return ProfileHandler{profileService: profileService}
}
