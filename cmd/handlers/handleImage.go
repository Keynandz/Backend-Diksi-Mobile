package handlers

import (
	"net/http"

	"golang/cmd/repositores"
	"github.com/labstack/echo/v4"
)
func GetImage(c echo.Context) error {
	images, err := repositories.GetImage()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, images)
}

func UploadImage(c echo.Context) error {
	
}