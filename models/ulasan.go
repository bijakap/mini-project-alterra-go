package models

import "gorm.io/gorm"

type Ulasan struct {
	gorm.Model
	Ulasan string `gorm:"not null" query:"ulasan" param:"ulasan" json:"ulasan" form:"ulasan"`
	Img string `query:"image" param:"image" json:"image" form:"image"`
	Id_museum int `query:"id_museum" param:"id_museum" json:"id_museum" form:"id_museum"`
	Id_User int `query:"id_User" param:"id_User" json:"id_User" form:"id_User"`
}