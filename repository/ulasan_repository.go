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
	Ulasan := models.Ulasan{Ulasan: ulasan, Id_User: id_user, Id_museum: id_museum, Img : image}
	if result := r.DB.Create(Ulasan); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repositoryPostgresLayerUlasan) GetUlasanByIdMuseum(id_museum int) []models.Ulasan {
	Ulasan := []models.Ulasan{}
	r.DB.Where("id_museum = ?", id_museum).Find(&Ulasan)
	return Ulasan
}

func NewRepositoryPostgresLayerUlasan(db *gorm.DB) domain.AdapterRepositoryUlasan {
	return &repositoryPostgresLayerUlasan{
		DB: db,
	}
}

