package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	UserID         uint         `gorm:"foreignKey:UserID"`
	Name           string       `gorm:"not null"`
	Description    string       `gorm:"not null"`
	PrepTime       string       `gorm:""`
	CookTime       string       `gorm:""`
	AdditionalTime string       `gorm:""`
	TotalTime      string       `gorm:""`
	Servings       string       `gorm:""`
	LikeCount      int          `gorm:"not null"`
	Instructions   string       `gorm:"not null"`
	Rating         int          `gorm:"not null;enum:1,2,3,4,5"`
	IsPublished    bool         `gorm:"not null;default:true"`
	Difficulty     string       `gorm:"not null;enum:easy,medium,hard"`
	Tags           string       `gorm:"not null"`
	Ingredients    []Ingredient `gorm:"foreignKey:RecipeID"`
}

type RecipeCreate struct {
	Name           string             `json:"name" form:"name" valid:"required,stringlength(3|50)"`
	Description    string             `json:"description" form:"description" valid:"required,stringlength(3|255)"`
	PrepTime       string             `json:"prep_time" form:"prep_time" valid:"stringlength(1|50)"`
	CookTime       string             `json:"cook_time" form:"cook_time" valid:"stringlength(1|50)"`
	AdditionalTime string             `json:"additional_time" form:"additional_time" valid:"stringlength(1|50)"`
	TotalTime      string             `json:"total_time" form:"total_time" valid:"stringlength(1|50)"`
	Servings       string             `json:"servings" form:"servings" valid:"stringlength(1|50)"`
	Instructions   string             `json:"instructions" form:"instructions" valid:"required,stringlength(10|255)"`
	IsPublished    bool               `json:"is_published" form:"is_published" valid:"required"`
	Difficulty     string             `json:"difficulty" form:"difficulty" valid:"required"`
	Tags           string             `json:"tags" form:"tags" valid:"stringlength(3|255)"`
	Ingredient     []IngredientCreate `json:"ingredients" form:"ingredients" valid:"required"`
}

type RecipeUpdate struct {
	Name           string `json:"name" form:"name" valid:"stringlength(3|50)"`
	Description    string `json:"description" form:"description" valid:"stringlength(3|255)"`
	PrepTime       string `json:"prep_time" form:"prep_time" valid:"stringlength(3|50)"`
	CookTime       string `json:"cook_time" form:"cook_time" valid:"stringlength(3|50)"`
	AdditionalTime string `json:"additional_time" form:"additional_time" valid:"stringlength(3|50)"`
	TotalTime      string `json:"total_time" form:"total_time" valid:"stringlength(3|50)"`
	Servings       string `json:"servings" form:"servings" valid:"stringlength(3|50)"`
	Instructions   string `json:"instructions" form:"instructions" valid:"stringlength(3|255)"`
	IsPublished    bool   `json:"is_published" form:"is_published"`
	Difficulty     string `json:"difficulty" form:"difficulty" valid:"enum(easy|medium|hard)"`
	Tags           string `json:"tags" form:"tags" valid:"stringlength(3|255)"`
}

type RecipeResource struct {
	ID             uint                 `json:"id"`
	UserID         uint                 `json:"user_id"`
	User           string               `json:"user"`
	Name           string               `json:"name"`
	Description    string               `json:"description"`
	PrepTime       string               `json:"prep_time"`
	CookTime       string               `json:"cook_time"`
	AdditionalTime string               `json:"additional_time"`
	TotalTime      string               `json:"total_time"`
	Servings       string               `json:"servings"`
	LikeCount      int                  `json:"like_count"`
	Instructions   string               `json:"instructions"`
	Rating         int                  `json:"rating"`
	IsPublished    bool                 `json:"is_published"`
	Difficulty     string               `json:"difficulty"`
	Tags           string               `json:"tags"`
	CreatedAt      string               `json:"created_at"`
	UpdatedAt      string               `json:"updated_at"`
	Ingredients    []IngredientResource `json:"ingredients"`
}
