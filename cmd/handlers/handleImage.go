package handlers

import (
	"net/http"

	"golang/cmd/models"
	"golang/cmd/repositories"
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
    var image models.Image
    if err := c.Bind(&image); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid request payload")
    }

    err := repositories.UploadImage(image.Name, image.Mading)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, "Image uploaded successfully")
}