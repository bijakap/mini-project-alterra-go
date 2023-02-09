package models

import "gorm.io/gorm"

type Ulasan struct {
	gorm.Model
	Ulasan string `gorm:"not null"`
	Img string
	Id_museum int 
	Id_User int
}