package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Id_museum int `json:"id_museum" form:"id_museum" `
	Img *string `gorm:"not null" json:"img" form:"img"`
}