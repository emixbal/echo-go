package controllers

import (
	"echo-go/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllPegawai(c echo.Context) error {
	result, err := models.FethAllPegawai()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func StorePegawai(c echo.Context) error {
	nama := c.FormValue("name")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telephone")

	result, err := models.StorePegawai(nama, alamat, telepon)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdatePegawai(c echo.Context) error {
	name := c.FormValue("name")
	alamat := c.FormValue("alamat")
	telephone := c.FormValue("telephone")
	id := c.Param("id")
	converted_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdatePegawai(converted_id, name, alamat, telephone)

	return c.JSON(http.StatusOK, result)
}

func DeletePegawai(c echo.Context) error {
	id := c.Param("id")
	converted_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeletePegawai(converted_id)

	return c.JSON(http.StatusOK, result)
}
