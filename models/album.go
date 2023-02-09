package models

import "gorm.io/gorm"

type Ablum struct {
	gorm.Model
	Id_museum int 
	Img string
}