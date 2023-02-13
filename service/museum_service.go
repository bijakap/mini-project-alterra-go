package service

import (
	"hi-story/domain"
	"hi-story/models"
)

type SvcMuseum struct {
	repo domain.AdapterRepositoryMuseum
}

func (s *SvcMuseum) CreateMuseumService(museum models.Museum) error {
	return nil
}

func (s *SvcMuseum) GetMuseumByIdService(id int) models.Museum {
	return models.Museum{}
}

func NewServiceMuseum(repo domain.AdapterRepositoryMuseum) domain.AdapterServiceMuseum {
	return &SvcMuseum {
		repo: repo,
	}
}