package repository

import (
	"hi-story/domain"
	"hi-story/models"

	"gorm.io/gorm"
)

type repositoryPostgresLayerUlasan struct {
	DB *gorm.DB
}

func (r *repositoryPostgresLayerUlasan) CreateUlasan(id_user, id_museum int, ulasan, image string) error {
	return nil
}

func (r *repositoryPostgresLayerUlasan) GetUlasanByIdMuseum(id_museum int) []models.Ulasan {
	return nil
}

func NewRepositoryPostgresLayerUlasan(db *gorm.DB) domain.AdapterRepositoryUlasan {
	return &repositoryPostgresLayerUlasan{
		DB: db,
	}
}

