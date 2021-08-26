package routes

import (
	"echo-go/controllers"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		fmt.Println("/ hitted")
		return c.String(http.StatusOK, "hi, im from echo")
	})

	e.GET("/pegawai", controllers.FetchAllPegawai)
	e.POST("/pegawai", controllers.StorePegawai)
	e.PUT("/pegawai/:id", controllers.UpdatePegawai)
	e.DELETE("/pegawai/:id", controllers.DeletePegawai)

	return e
}
