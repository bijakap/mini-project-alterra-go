package config

import (
	"fmt"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/subosito/gotenv"

	m "hi-story/models"
)

var DB *gorm.DB

func init() {
	gotenv.Load()
}

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}

func InitMigrate(){
	DB.AutoMigrate(&m.User{})
	DB.AutoMigrate(&m.Museum{})
	DB.AutoMigrate(&m.Ulasan{})

	admin := &m.User{Nama: "Admin", Email: "admin@gmail.com", Password: "admin123", Role : "Admin"}

	DB.Create(admin)
}