package controller

import (
	"hi-story/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EchoControllerAlbum struct {
	Svc domain.AdapterServiceAlbum
}

func (ec *EchoControllerAlbum) CreateAlbumController(c echo.Context) error {
	return c.String(http.StatusOK, "Create Album")
}

func (ec *EchoControllerAlbum) GetAlbumByIdController(c echo.Context) error {
	return c.String(http.StatusOK, "Get Album By ID")
}