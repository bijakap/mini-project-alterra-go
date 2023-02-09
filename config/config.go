package config

import (
	"fmt"
	"os"

	"github.com/subosito/gotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

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
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err.Error())
	}

	return DB
}

func InitMigrate(){
	DB.AutoMigrate(&m.User{}, &m.Museum{}, &m.Ulasan{}, &m.Ablum{})
	
	if res := DB.Find(&m.User{}); res.RowsAffected == 0 {
		admin := &m.User{Nama: "Admin", Email: "admin@gmail.com", Password: "admin123", Role : "Admin"}
		DB.Create(admin)
		firstUser := &m.User{Nama: "Bijak Algifan Putra", Email: "bijak.algifan.p@gmail.com", Password: "qwerty123"}
		DB.Create(firstUser)
	}

	if resMuseum := DB.Find(&m.Museum{}); resMuseum.RowsAffected == 0{
		museum := &m.Museum{Nama: "Konferensi Asia Afrika", Deskripsi: "Ini Test", Alamat: "Jln Asia Afrika"}
		DB.Create(museum)
	}

	if resUlasan := DB.Find(&m.Ulasan{}); resUlasan.RowsAffected == 0 {
		ulasan := &m.Ulasan{Ulasan: "test ulasan", Id_User: 2, Id_museum: 1}
		DB.Create(ulasan)
	}

}