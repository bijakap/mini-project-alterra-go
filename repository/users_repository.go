package repository

import (
	"fmt"

	"gorm.io/gorm"

	"hi-story/domain"
	"hi-story/models"
)

type repositoryPostgresLayer struct {
	DB *gorm.DB
}

func (r *repositoryPostgresLayer) CreateUsers(user models.User) error {
	res := r.DB.Create(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repositoryPostgresLayer) GetAll() []models.User {
	users := []models.User{}
	r.DB.Find(&users)
	return users
}

func NewPostgresRepository(db *gorm.DB) domain.AdapterRepository {
	return &repositoryPostgresLayer{
		DB: db,
	}
}