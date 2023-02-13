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
	
	if id == 0 {
		return c.String(http.StatusBadRequest, "Kesalahan dalam id")
	}
	
	message := fmt.Sprintf("Test API Ulasan GetUlasanByIDMuseum %d", id)
	ulasan := ec.Svc.GetUlasanByIdMuseumService(id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : message,
		"ulasan" : ulasan,
	})
}

func (ec *EchoControllerUlasan) CreateUlasanController(c echo.Context) error {
	// text, image := c.QueryParam("ulasan"), c.QueryParam("gambar")
	// id_user, _ := strconv.Atoi(c.QueryParam("id_user"))
	// id_museum, _ := strconv.Atoi(c.QueryParam("id_museum"))

	ulasan := models.Ulasan{}
	c.Bind(&ulasan)

	// if c.QueryParam("ulasan") == "" &&  ulasan.Ulasan == ""{
	// 	return c.String(http.StatusBadRequest, "Ulasan Kosong")
	// }

	if ulasan.Ulasan == "" || ulasan.Id_User == 0 || ulasan.Id_museum == 0{
		return c.String(http.StatusBadRequest, "Ulasan Kosong atau terdapat kesalahan")
	}
	
	err := ec.Svc.CreateUlasanService(ulasan.Id_User,ulasan.Id_museum,ulasan.Ulasan,ulasan.Img)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message" : err.Error(),
		})
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "Ulasan Telah ditambahkan",
		"ulasan" : ulasan,
	})
}
