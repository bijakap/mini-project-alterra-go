package models

import "gorm.io/gorm"

type Museum struct {
	gorm.Model
	Nama *string `gorm:"not null" query:"nama" param:"nama" json:"nama" form:"nama"`
	Deskripsi *string `gorm:"not null" query:"desc" param:"desc" json:"desc" form:"desc"`
	Alamat *string  `gorm:"not null" query:"alamat" param:"alamat" json:"alamat" form:"alamat"`
	Jadwal *string `query:"jadwal" param:"jadwal" json:"jadwal" form:"jadwal"`
	Kontak *string `query:"kontak" param:"kontak" json:"kontak" form:"kontak"`
	Gambar *string `query:"image" param:"image" json:"image" form:"image"`
	Ablum []Ablum `gorm:"foreignKey:Id_museum;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Ulasan []Ulasan `gorm:"foreignKey:Id_museum;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}