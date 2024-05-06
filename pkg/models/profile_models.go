package models

type Profile struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserID   uint   `json:"user_id" gorm:"forignKey:ID;unique;not null"`
	FullName string `json:"full_name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Gender   string `json:"gender" gorm:"not null;enum(M|F)"`
	Address  string `json:"address" gorm:"not null"`
	Photo    string `json:"photo"`
}

type ProfileCreate struct {
	FullName string `json:"full_name" form:"full_name" valid:"required,stringlength(3|50)"`
	Email    string `json:"email" form:"email" valid:"required,email"`
	Gender   string `json:"gender" form:"gender" valid:"required,in(M|F)"`
	Address  string `json:"address" form:"address" valid:"required,stringlength(3|50)"`
	Photo    string `json:"photo" form:"photo"`
}

type ProfileUpdate struct {
	FullName string `json:"full_name" form:"full_name" valid:"stringlength(3|50)"`
	Email    string `json:"email" form:"email" valid:"email"`
	Gender   string `json:"gender" form:"gender" valid:"in(M|F)"`
	Address  string `json:"address" form:"address" valid:"stringlength(3|50)"`
	Photo    string `json:"photo" form:"email" valid:"stringlength(3|50)"`
}

type ProfileResource struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	Usermame string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Address  string `json:"address"`
	Photo    string `json:"photo"`
}
