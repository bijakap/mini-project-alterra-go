package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Role string `gorm:"not null;default:user"`
	Email string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Nama string `gorm:"not null"`
	Profile_pic string
	Remember_me string
}

