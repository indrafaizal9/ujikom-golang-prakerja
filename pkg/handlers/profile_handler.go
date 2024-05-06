package handlers

import (
	"strconv"
	"ujikom/pkg/helpers"
	"ujikom/pkg/models"
	"ujikom/pkg/services"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileService services.ProfileService
}

func NewProfileHandler(profileService services.ProfileService) ProfileHandler {
	return ProfileHandler{profileService: profileService}
}

func (h *ProfileHandler) GetAllProfiles(c *gin.Context) {
	h.profileService.GetAllProfiles(c)
}

func (h *ProfileHandler) GetMyProfile(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	h.profileService.GetProfileByUserID(c, user.ID)
}

func (h *ProfileHandler) CreateMyProfile(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	h.profileService.CreateProfile(c, user.ID)
}

func (h *ProfileHandler) UpdateMyProfile(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	h.profileService.UpdateProfile(c, user.ID)
}

func (h *ProfileHandler) GetAllCollections(c *gin.Context) {
	h.profileService.GetAllCollections(c)
}

func (h *ProfileHandler) GetMyRecipes(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	h.profileService.GetRecipesByUserID(c, user.ID)
}

func (h *ProfileHandler) GetMyCollections(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	h.profileService.GetMyCollections(c, user.ID)
}

func (h *ProfileHandler) GetCollection(c *gin.Context) {
	user := c.MustGet("user").(models.User).ID
	parseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResBadRequest(c, "Invalid ID")
		return
	}
	h.profileService.GetCollection(c, user, uint(parseID))
}

func (h *ProfileHandler) CreateCollection(c *gin.Context) {
	userID := c.MustGet("user").(models.User).ID
	request := models.CollectionCreate{}
	err := c.Bind(&request)
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}
	h.profileService.CreateCollection(c, userID, request)
}

func (h *ProfileHandler) UpdateCollection(c *gin.Context) {
	userID := c.MustGet("user").(models.User).ID
	parseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResBadRequest(c, "Invalid ID")
		return
	}
	request := models.CollectionUpdate{}
	err = c.Bind(&request)
	if err != nil {
		helpers.ResBadRequest(c, err.Error())
		return
	}
	h.profileService.UpdateCollection(c, userID, uint(parseID), request)
}

func (h *ProfileHandler) DeleteCollection(c *gin.Context) {
	userID := c.MustGet("user").(models.User).ID
	parseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResBadRequest(c, "Invalid ID")
		return
	}
	h.profileService.DeleteCollection(c, userID, uint(parseID))
}
