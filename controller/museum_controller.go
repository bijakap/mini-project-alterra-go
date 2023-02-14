package controller

import (
	"fmt"
	"hi-story/domain"
	"hi-story/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EchoControlleMuseum struct {
	Svc domain.AdapterServiceMuseum
}

func (ec *EchoControlleMuseum) CreateMuseumController(c echo.Context) error {
	museum := models.Museum{}
	c.Bind(&museum)

	err := ec.Svc.CreateMuseumService(museum)
	if err != nil {
		return c.String(http.StatusInternalServerError, "museum error")
	}

	// return c.String(http.StatusOK, "create museum")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "create museum",
		"data" : museum,
	})
}

func (ec *EchoControlleMuseum) GetMuseumByIdController(c echo.Context) error {
	id,_ := strconv.Atoi(c.Param("id"))
	message := fmt.Sprintf("get museum by id = %d", id)

	museum := ec.Svc.GetMuseumByIdService(id)
	
	// return c.String(http.StatusOK, message)
	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"message" : message,
		"data" : museum,
	}, "")
}

