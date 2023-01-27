package routes

import (
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
	"net/http"
)

func New() *echo.Echo{
	e := echo.New()

	//Routing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	return e
}