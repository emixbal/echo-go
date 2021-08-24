package routes

import (
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

	return e
}
