package controller

import (
	"golang/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DataBarang(c echo.Context) error {

	result, err := models.DataBarang()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}
