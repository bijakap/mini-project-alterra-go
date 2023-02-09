package routes

import (
	"hi-story/config"
	"hi-story/controller"
	"hi-story/repository"
	"hi-story/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterUlasanGroup(e *echo.Echo){
	DB := config.InitDB()
	repo := repository.NewRepositoryPostgresLayerUlasan(DB)
	service := service.NewServiceUlasan(repo)
	cont := controller.EchoControllerUlasan{Svc: service}

	apiUlasan := e.Group("/ulasan", middleware.Logger())

	apiUlasan.GET("", cont.GetUlasanByIdMuseumController)
	apiUlasan.POST("", cont.CreateUlasanController)
	
}