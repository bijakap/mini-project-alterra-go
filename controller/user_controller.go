package controller

import (
	"hi-story/domain"
	"hi-story/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EchoController struct {
	Svc domain.AdapterService
}

func (ec *EchoController) CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	err := ec.Svc.CreateUserService(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message" : err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "success",
		"user" : user,
	})
}

func (ec *EchoController) GetAllUsersController(c echo.Context) error {
	users := ec.Svc.GetAllUsersService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"Message" : "get all users",
		"users" : users,
	}, " ")
}

func (ec *EchoController) AuthUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	token, statusCode := "", 400
	if (user.Email != "" && user.Password != ""){
		token, statusCode = ec.Svc.AuthUser(user.Email, user.Password)
	} else if (c.QueryParam("email") != "" && c.QueryParam("password") != ""){
		token, statusCode = ec.Svc.AuthUser(c.QueryParam("email"), c.QueryParam("password"))
	} else {
		return c.JSONPretty(http.StatusBadRequest, map[string]interface{}{
			"Message" : "400 - Bad Request",
		}, " ")
	}
	
	switch (statusCode) {
	case http.StatusInternalServerError : 
		return c.JSONPretty(http.StatusInternalServerError, map[string]interface{}{
			"Message" : "500 - Servel Internal Error",
			"data" : user,
			"email" : c.QueryParam("email"),
			"password" : c.QueryParam("password"),
		}, " ")
	case http.StatusUnauthorized : 
			return c.JSONPretty(http.StatusUnauthorized, map[string]interface{}{
				"Message" : "403 - Unauthorized",
				"data" : user,
				"email" : c.QueryParam("email"),
				"password" : c.QueryParam("password"),
			}, " ")
	case http.StatusBadRequest :
		return c.JSONPretty(http.StatusBadRequest, map[string]interface{}{
			"Message" : "400 - Bad Request",
			"data" : user,
			"email" : c.QueryParam("email"),
			"password" : c.QueryParam("password"),
		}, " ")
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"message" : "Success",
		"token" : token,
		"binding" : user,
		"email" : c.QueryParam("email"),
		"password" : c.QueryParam("password"),
	}, " ")
}

