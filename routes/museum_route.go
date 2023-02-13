package routes

import (
	"hi-story/config"
	"hi-story/controller"
	"hi-story/repository"
	"hi-story/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterMuseumGroup(e *echo.Echo){
	DB := config.InitDB()
	repo := repository.NewRepositoryPostgresLayerMuseum(DB)
	svc := service.NewServiceMuseum(repo)
	cont := controller.EchoControlleMuseum{
		Svc: svc,
	}

	apiMuseum := e.Group("/museum", middleware.Logger())

	apiMuseum.GET("/:id", cont.GetMuseumByIdController)
	apiMuseum.POST("", cont.CreateMuseumController)


}