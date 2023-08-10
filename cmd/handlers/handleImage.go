package handlers

import (
	"fmt"
	"io"

	"net/http"

	"golang/cmd/repositories"

	"github.com/labstack/echo/v4"
)

func GetImageByID(c echo.Context) error {
    imageID := c.Param("id")
    image, err := repositories.GetImageByID(imageID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    c.Response().Header().Set("Content-Type", "image/png")
    c.Response().Write(image.Mading)
	fmt.Println("Gambar Berhasil Dicomot")
    return nil
}

func UploadImage(c echo.Context) error {
    file, err := c.FormFile("mading")
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Failed to get image file")
    }

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to open uploaded file")
	}
	defer src.Close()
	
	imageData, err := io.ReadAll(src)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to read uploaded file")
	}

    imageName := c.FormValue("name")

    err = repositories.UploadImage(imageName, imageData)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, "Image uploaded successfully")
}