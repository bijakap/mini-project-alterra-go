package service

import (
	"hi-story/domain"
	"hi-story/models"
)

type svcUlasan struct {
	repo domain.AdapterRepositoryUlasan
}

func (s *svcUlasan) CreateUlasanService(id_user, id_museum int, ulasan, image string) error {
	return s.repo.CreateUlasan(id_user, id_museum, ulasan, image)
}

func (s *svcUlasan) GetUlasanByIdMuseumService(id_museum int) []models.Ulasan {
	return s.repo.GetUlasanByIdMuseum(id_museum)
}

func NewServiceUlasan(repo domain.AdapterRepositoryUlasan) domain.AdapterServiceUlasan {
	return &svcUlasan {
		repo: repo,
	}
}