package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null" `
	Password string `json:"password" gorm:"not null"`
	Role     string `json:"role" gorm:"enum:admin,user"`
	IsActive bool   `json:"is_active" gorm:"default:true"`
}

type UserCreate struct {
	Username string `json:"username" form:"username" valid:"required,stringlength(3|20)"`
	Password string `json:"password" form:"password" valid:"required,stringlength(6|20)"`
	Role     string `json:"role" form:"role" valid:"required,in(admin|user)"`
}

type UserResource struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
}

type UserUpdate struct {
	Username string  `json:"username" valid:"stringlength(3|20)"`
	Password string `json:"password" valid:"stringlength(6|20)"`
	Role     string `json:"role" valid:"in(admin|user)"`
	IsActive bool   `json:"is_active"`
}

type Login struct {
	Username string `json:"username" binding:"required" valid:"stringlength(3|20)"`
	Password string `json:"password" binding:"required" valid:"stringlength(6|20)"`
}
