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

	// binding
	if (user.Nama != "" || user.Email != "" || user.Password != ""){ 
		err := ec.Svc.CreateUserService(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message" : err.Error(),
			})
		}
	// Query Param
	} else if (c.QueryParam("nama") != "" || c.QueryParam("email") != "" || c.QueryParam("password") != ""){ 
		err := ec.Svc.CreateUserService(models.User{Nama: c.QueryParam("nama"), Email: c.QueryParam("email"), Password: c.QueryParam("password"),})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message" : err.Error(),
			})
		}
	} else {
		return c.String(http.StatusBadRequest, "Bad Request")
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
		token, statusCode, user = ec.Svc.AuthUser(user.Email, user.Password)
	} else if (c.QueryParam("email") != "" && c.QueryParam("password") != ""){
		token, statusCode, user = ec.Svc.AuthUser(c.QueryParam("email"), c.QueryParam("password"))
	} else {
		return c.JSONPretty(http.StatusBadRequest, map[string]interface{}{
			"Message" : "400 - Bad Request",
		}, " ")
	}
	
	switch (statusCode) {
	case http.StatusInternalServerError : 
		return c.JSONPretty(http.StatusInternalServerError, map[string]interface{}{
			"Message" : "500 - Servel Internal Error",
		}, " ")
	case http.StatusUnauthorized : 
			return c.JSONPretty(http.StatusUnauthorized, map[string]interface{}{
				"Message" : "403 - Unauthorized",
			}, " ")
	case http.StatusBadRequest :
		return c.JSONPretty(http.StatusBadRequest, map[string]interface{}{
			"Message" : "400 - Bad Request",
		}, " ")
	case http.StatusNotFound :
		return c.JSONPretty(http.StatusNotFound, map[string]interface{}{
			"Message" : "404 - Not Found",
		}, " ")
	}

	userRespone := map[string]interface{}{
		"nama" : user.Nama,
		"email" : user.Email,
		"profile_image" : user.Profile_pic,
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"message" : "Success",
		"token" : token,
		"data" : userRespone,
	}, " ")
}

