package models

import "gorm.io/gorm"

type Museum struct {
	gorm.Model
	Nama *string `gorm:"not null"`
	Deskripsi *string `gorm:"not null"`
	Alamat *string  `gorm:"not null"`
	Jadwal *string 
	Kontak *string
	Gambar *string 
	Ablum []Ablum `gorm:"foreignKey:Id_museum;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Ulasan []Ulasan `gorm:"foreignKey:Id_museum;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}