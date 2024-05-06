package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	UserID   uint   `gorm:"foreignKey:ID"`
	RecipeID uint   `gorm:"foreignKey:ID"`
	Rating   int    `gorm:"not null;enum:1,2,3,4,5"`
	Review   string `gorm:"not null"`
}

type ReviewCreate struct {
	UserID   uint   `json:"user_id" form:"user_id" valid:"required,numeric"`
	RecipeID uint   `json:"recipe_id" form:"recipe_id" valid:"required,numeric"`
	Rating   int    `json:"rating" form:"rating" valid:"required,range(1|5)"`
	Review   string `json:"review" form:"review" valid:"required,stringlength(10|255)"`
}

type ReviewUpdate struct {
	Rating int    `json:"rating" form:"rating" valid:"range(1|5)"`
	Review string `json:"review" form:"review" valid:"stringlength(10|255)"`
}

type ReviewResource struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	RecipeID uint   `json:"recipe_id"`
	Rating   int    `json:"rating"`
	Review   string `json:"review"`
}
