package database

import (
	"fmt"
	"log"
	"ujikom/pkg/models"
)

type Seeder struct {
}

func (s *Seeder) SeederTagsAndLabels(){
	tags := []models.TagsAndLabels{
		{Name: "Breakfast", Type: "tag"},
		{Name: "Lunch", Type: "tag"},
		{Name: "Dinner", Type: "tag"},
		{Name: "Snack", Type: "tag"},
		{Name: "Dessert", Type: "tag"},
		{Name: "Vegetarian", Type: "tag"},
		{Name: "Vegan", Type: "tag"},
		{Name: "Gluten Free", Type: "tag"},
		{Name: "Keto", Type: "tag"},
		{Name: "Low Carb", Type: "tag"},
	}

	labels := []models.TagsAndLabels{
		{Name: "Healthy", Type: "label"},
		{Name: "Quick", Type: "label"},
		{Name: "Easy", Type: "label"},
		{Name: "Spicy", Type: "label"},
		{Name: "Sweet", Type: "label"},
		{Name: "Sour", Type: "label"},
		{Name: "Salty", Type: "label"},
		{Name: "Bitter", Type: "label"},
		{Name: "Savoury", Type: "label"},
	}

	err := DB.Create(&tags)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Create(&labels)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tags and Labels Seeded!")
}

func (s *Seeder) SeederUsers(){
	users := []models.User{}

	err := DB.Create(&users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Users Seeded!")
}