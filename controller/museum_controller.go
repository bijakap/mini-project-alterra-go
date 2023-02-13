package controller

import (
	"hi-story/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EchoControlleMuseum struct {
	Svc domain.AdapterServiceMuseum
}

func (ec *EchoControlleMuseum) CreateMuseumController(c echo.Context) error {
	return c.String(http.StatusOK, "create museum")
}

func (ec *EchoControlleMuseum) GetMuseumByIdController(c echo.Context) error {
	return c.String(http.StatusOK, "get museum by id")
}

