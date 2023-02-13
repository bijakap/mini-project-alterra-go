package repository

import (
	"hi-story/domain"
	"hi-story/models"

	"gorm.io/gorm"
)

type repositoryPostgresLayerMuseum struct {
	DB *gorm.DB
}

func (r *repositoryPostgresLayerMuseum) CreateMuseum(museum models.Museum) error {
	return nil
}

func (r *repositoryPostgresLayerMuseum) GetMuseumById(id int) models.Museum {
	return models.Museum{}
}

func NewRepositoryPostgresLayerMuseum(db *gorm.DB) domain.AdapterRepositoryMuseum {
	return &repositoryPostgresLayerMuseum{
		DB: db,
	}
}
