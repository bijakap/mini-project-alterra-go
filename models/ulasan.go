package models

import "gorm.io/gorm"

type Ulasan struct {
	gorm.Model
	Ulasan string `gorm:"not null"`
	Img string `gorm:"not null"`
	Id_museum int 
}