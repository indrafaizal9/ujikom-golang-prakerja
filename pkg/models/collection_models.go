package models

type Collection struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Public      bool   `json:"public"`
	UserID      int    `json:"user_id"`
	Description string `json:"description"`
}

type CollectionCreate struct {
	Name        string `json:"name" form:"name" valid:"required,string,length(1|255)"`
	Public      bool   `json:"public" form:"public" valid:"required,bool"`
	Description string `json:"description" form:"description" valid:"required,string,length(1|255)"`
}

type CollectionUpdate struct {
	Name        string `json:"name" form:"name" valid:"string,length(1|255)"`
	Public      bool   `json:"public" form:"public" valid:"bool"`
	Description string `json:"description" form:"description" valid:"string,length(1|255)"`
}

type CollectionRecipesPivot struct {
	CollectionID int `json:"collection_id" gorm:"foreignkey:ID;association_foreignkey:ID"`
	RecipeID     int `json:"recipe_id" gorm:"foreignkey:ID;association_foreignkey:ID"`
}
