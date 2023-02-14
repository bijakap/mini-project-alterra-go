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
	if result := r.DB.Create(&museum); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repositoryPostgresLayerMuseum) GetMuseumById(id int) models.Museum {
	museum := models.Museum{}
	r.DB.Where("id = ?", id).Find(&museum)
	return museum
}

func NewRepositoryPostgresLayerMuseum(db *gorm.DB) domain.AdapterRepositoryMuseum {
	return &repositoryPostgresLayerMuseum{
		DB: db,
	}
}
