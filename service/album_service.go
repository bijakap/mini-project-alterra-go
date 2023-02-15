package service

import (
	"hi-story/domain"
	"hi-story/models"
)

type svcAlbum struct {
	repo domain.AdapterRepositoryAlbum
}

func (s *svcAlbum) CreateAlbumService(album models.Album) error {
	return s.repo.CreateAlbum(album)
}

func (s *svcAlbum) GetAlbumByIdMuseumService(id int) []models.Album {
	return s.repo.GetAlbumByIdMuseum(id)
}

func NewServiceAlbum(repo domain.AdapterRepositoryAlbum) domain.AdapterServiceAlbum {
	return &svcAlbum {
		repo: repo,
	}
}