package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Role string `gorm:"not null;default:user" json:"role" form:"role" `
	Email string `gorm:"unique;not null" query:"email" param:"email" json:"email" form:"email" `
	Password string `gorm:"not null" query:"password" param:"password" json:"password" form:"password" `
	Nama string `gorm:"not null" json:"nama" form:"nama" `
	Profile_pic *string `json:"profile_pic" form:"profile_pic" `
	Remember_me *string `json:"remember_me" form:"remember_me" `
	Ulasan []Ulasan `gorm:"foreignKey:Id_User;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

