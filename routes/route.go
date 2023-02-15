package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// "github.com/labstack/echo/v4/middleware"
	"net/http"
)

func New() *echo.Echo{
	e := echo.New()
	e.Use(middleware.Logger())

	//Routing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.GET("test",func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	RegisterUserGroup(e)
	RegisterUlasanGroup(e)
	RegisterMuseumGroup(e)
	RegisterAlbumGroup(e)

	return e
}