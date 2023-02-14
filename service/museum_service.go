package service

import (
	"hi-story/domain"
	"hi-story/models"
)

type SvcMuseum struct {
	repo domain.AdapterRepositoryMuseum
}

func (s *SvcMuseum) CreateMuseumService(museum models.Museum) error {
	return s.repo.CreateMuseum(museum)
}

func (s *SvcMuseum) GetMuseumByIdService(id int) models.Museum {
	return s.repo.GetMuseumById(id)
}

func NewServiceMuseum(repo domain.AdapterRepositoryMuseum) domain.AdapterServiceMuseum {
	return &SvcMuseum {
		repo: repo,
	}
}