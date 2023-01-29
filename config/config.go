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
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return DB
}

func InitMigrate(){
	DB.AutoMigrate(&m.User{})
	DB.AutoMigrate(&m.Museum{})
	DB.AutoMigrate(&m.Ulasan{})

	

	res := DB.Find(&m.User{})
	if(res.RowsAffected == 0){
		admin := &m.User{Nama: "Admin", Email: "admin@gmail.com", Password: "admin123", Role : "Admin"}
		DB.Create(admin)
	}
	

}