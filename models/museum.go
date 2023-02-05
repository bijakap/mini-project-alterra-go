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
	// Ablum []Ablum `gorm:"foreignKey:Id_museum"`
	// Ulasan []Ulasan `gorm:"foreignKey:Id_museum"`
}