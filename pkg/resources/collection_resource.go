package resources

import "ujikom/pkg/models"

func CollectionMake(collection models.Collection, collectionResource *models.CollectionResource) {
	collectionResource.ID = collection.ID
	collectionResource.UserID = collection.UserID
	collectionResource.Name = collection.Name
	collectionResource.Description = collection.Description
}

func CollectionCollection(collections []models.Collection, collectionResources *[]models.CollectionResource) {
	for _, collection := range collections {
		collectionResource := models.CollectionResource{}
		CollectionMake(collection, &collectionResource)
		*collectionResources = append(*collectionResources, collectionResource)
	}
}