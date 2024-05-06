package models

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Amount      string `gorm:"not null"`
	Unit        string `gorm:"not null"`
	RecipeID    uint   `gorm:"foreignKey:RecipeID"`
}

type IngredientCreate struct {
	Name        string `json:"name" form:"name" valid:"required,stringlength(3|50)"`
	Description string `json:"description" form:"description" valid:"required,stringlength(1|255)"`
	Amount      string `json:"amount" form:"amount" valid:"required,stringlength(1|50)"`
	Unit        string `json:"unit" form:"unit" valid:"required,stringlength(1|50)"`
}

type IngredientUpdate struct {
	Name        string `json:"name" form:"name" valid:"stringlength(1|50)"`
	Description string `json:"description" form:"description" valid:"stringlength(1|255)"`
	Amount      string `json:"amount" form:"amount" valid:"stringlength(1|50)"`
	Unit        string `json:"unit" form:"unit" valid:"stringlength(1|50)"`
}

type IngredientResource struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      string `json:"amount"`
	Unit        string `json:"unit"`
}
