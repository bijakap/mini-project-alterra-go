package controller

import (
	"fmt"
	"hi-story/domain"
	"hi-story/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EchoControllerUlasan struct {
	Svc domain.AdapterServiceUlasan
}

func (ec *EchoControllerUlasan) GetUlasanByIdMuseumController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	message := fmt.Sprintf("Test API Ulasan GetUlasanByIDMuseum %d", id)
	ulasan := ec.Svc.GetUlasanByIdMuseumService(id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : message,
		"ulasan" : ulasan,
	})
}

func (ec *EchoControllerUlasan) CreateUlasanController(c echo.Context) error {
	text, image,id_user,id_museum := c.QueryParam("ulasan"), c.QueryParam("gambar"), c.QueryParam("id_user"), c.QueryParam("id_museum")

	ulasan := models.Ulasan{}
	c.Bind(&ulasan)

	if c.QueryParam("ulasan") == ""{
		return c.String(http.StatusBadRequest, "Ulasan Kosong")
	}

	err := ec.Svc.CreateUlasanService(2,1,text,image)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message" : err.Error(),
		})
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"text" : text,
		"image" : image,
		"param_id_user" : id_user,
		"param_id_museum" : id_museum,
		"ulasan" : ulasan,
	})
}
