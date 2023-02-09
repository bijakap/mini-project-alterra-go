package service

import (
	"hi-story/domain"
	"hi-story/models"
)

type svcUlasan struct {
	repo domain.AdapterRepositoryUlasan
}

func (s *svcUlasan) CreateUlasanService(id_user, id_museum int, ulasan, image string) error {
	return nil
}

func (s *svcUlasan) GetUlasanByIdMuseumService() []models.Ulasan {
	return nil
}

func NewServiceUlasan(repo domain.AdapterRepositoryUlasan) domain.AdapterServiceUlasan {
	return &svcUlasan {
		repo: repo,
	}
}