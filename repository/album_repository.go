package repository

import (
	"hi-story/domain"
	"hi-story/models"

	"gorm.io/gorm"
)

type repositoryPostgresLayerAlbum struct {
	DB *gorm.DB
}

func (r *repositoryPostgresLayerAlbum) CreateAlbum(album models.Album) error {
	if result := r.DB.Create(&album); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repositoryPostgresLayerAlbum) GetAlbumByIdMuseum(id int) []models.Album {
	album := []models.Album{}
	r.DB.Where("id_museum = ?", id).Find(&album)
	return album
}

func NewRepositoryPostgresLayerAlbum(db *gorm.DB) domain.AdapterRepositoryAlbum {
	return &repositoryPostgresLayerAlbum {
		DB : db,
	}
}