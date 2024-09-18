package routes

import (
	"golang/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Selamat Datang Chiyogami")
	})

	e.GET("/barang", controller.DataBarang)
	return e

}
