package routes

import (
	"hi-story/config"
	"hi-story/controller"
	"hi-story/repository"
	"hi-story/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterAlbumGroup(e *echo.Echo) {
	DB := config.InitDB()
	repo := repository.NewRepositoryPostgresLayerAlbum(DB)
	svc := service.NewServiceAlbum(repo)
	cont := controller.EchoControllerAlbum{
		Svc: svc,
	}

	apiAlbum := e.Group("/album", middleware.Logger())

	apiAlbum.GET("/:id", cont.GetAlbumByIdController)
	apiAlbum.POST("", cont.CreateAlbumController)	
}