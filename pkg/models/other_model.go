package models

type Likes struct {
	UserID   uint `json:"user_id" gorm:"foreignKey:ID"`
	RecipeID uint `json:"recipe_id" gorm:"foreignKey:ID"`
}

type TagsAndLabels struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Type string `json:"type" gorm:"enum:tag,label"`
}

type MealType struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Type        string `json:"type"`
	IsPublished bool   `json:"is_published"`
}
