package config

import (
	"fmt"
	"os"

	"github.com/subosito/gotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	con "hi-story/constants"
	m "hi-story/models"
)

var DB *gorm.DB

func init() {
	gotenv.Load()
}

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err.Error())
	}

	return DB
}

func InitMigrate(){
	DB.AutoMigrate(&m.User{}, &m.Museum{}, &m.Ulasan{}, &m.Album{})
	
	if res := DB.Find(&m.User{}); res.RowsAffected == 0 {
		profile_image := con.DEFAULT_PROFILE_PIC

		admin := &m.User{Nama: "Admin", Email: "admin@gmail.com", Password: "admin123", Role : "Admin"}
		admin.Profile_pic = &profile_image

		DB.Create(admin)
		firstUser := &m.User{Nama: "Bijak Algifan Putra", Email: "bijak.algifan.p@gmail.com", Password: "qwerty123"}
		DB.Create(firstUser)
	}

	if resMuseum := DB.Find(&m.Museum{}); resMuseum.RowsAffected == 0{
		// museum := &m.Museum{Nama: "Konferensi Asia Afrika", Deskripsi: "Ini Test", Alamat: "Jln Asia Afrika"}
		museum := &m.Museum{}
		Nama := "Konferensi Asia Afrika"
		Deskripsi := "Konferensi Asia Afrika"
		Alamat := "Jln Asia Afrika"
		default_image := con.DEFAULT_IMAGE

		museum.Nama = &Nama
		museum.Deskripsi = &Deskripsi
		museum.Alamat = &Alamat
		museum.Gambar = &default_image

		DB.Create(museum)
		Nama = "KAA kedua"
		museum.ID = 2
		DB.Create(museum)
 	}

	if resUlasan := DB.Find(&m.Ulasan{}); resUlasan.RowsAffected == 0 {
		ulasan := &m.Ulasan{Ulasan: "test ulasan", Id_User: 2, Id_museum: 1}
		DB.Create(ulasan)
	}

	if resAlbum := DB.Find(&m.Album{}); resAlbum.RowsAffected == 0 {
		default_image := con.DEFAULT_IMAGE
		album1 := &m.Album{Id_museum: 1}
		album2 := &m.Album{Id_museum: 1}
		album3 := &m.Album{Id_museum: 2}
		album1.Img = &default_image
		album2.Img = &default_image 
		album3.Img = &default_image
		DB.Create(album1)
		DB.Create(album2)
		DB.Create(album3)
	}

}