package models

import "gorm.io/gorm"

type Museum struct {
	gorm.Model
	Nama string `gorm:"not null"`
	Deskripsi string `gorm:"not null"`
	Alamat string  `gorm:"not null"`
	Jadwal string `gorm:"not null"`
	Kontak string `gorm:"not null"`
	Gambar string `gorm:"not null"`
	Ulasan []Ulasan `gorm:"Foreignkey:Id_museum;association_foreignkey:ID"`
}