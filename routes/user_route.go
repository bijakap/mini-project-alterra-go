package routes

import (
	"hi-story/config"
	"hi-story/controller"
	"hi-story/repository"
	"hi-story/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterUserGroup(e *echo.Echo) {
	DB := config.InitDB()
	repo := repository.NewPostgresRepository(DB)
	svc := service.NewServiceUser(repo)

	cont := controller.EchoController{Svc : svc}

	// apiUser := e.Group("/users", 
	// 	middleware.Logger(),
	// )

	e.POST("/auth", cont.AuthUserController)

	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Test API user")
	})
	e.GET("/testall", cont.GetAllUsersController)
	
	
	// gotenv.Load()
	// apiUser.Use(middleware.JWT([]byte(os.Getenv("API_SECRET"))))
	// apiUser.GET("", cont.GetAllUsersController, middleware.JWTWithConfig(
	// 	middleware.JWTConfig{
	// 		SigningKey: []byte(os.Getenv("API_SECRET")),
	// 		ErrorHandlerWithContext: func(err error, c echo.Context) error {
	// 			return c.JSONPretty(404, map[string]interface{}{
	// 				"messages": "token invalid",
	// 			}, "  ")
	// 		},
	// 		SuccessHandler: func(c echo.Context) {
	// 		},
	// 	},
	// ))
	// apiUser.POST("", cont.CreateUserController)

}