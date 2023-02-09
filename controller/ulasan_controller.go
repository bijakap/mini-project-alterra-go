package controller

import (
	"hi-story/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EchoControllerUlasan struct {
	Svc domain.AdapterServiceUlasan
}

func (ec *EchoControllerUlasan) CreateUlasanController(c echo.Context) error {
	return c.String(http.StatusOK, "Test API Ulasan CreateUlasan")
}

func (ec *EchoControllerUlasan) GetUlasanByIdMuseumController(c echo.Context) error {
	return c.String(http.StatusOK, "Test API Ulasan GetUlasanByIDMuseum")
}