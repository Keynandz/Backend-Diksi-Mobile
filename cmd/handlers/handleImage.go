package handlers

import (
	"fmt"
	"io"
	"strconv"

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

func GetImageByTimestampOrder(c echo.Context) error {
	order, err := strconv.Atoi(c.Param("order"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid image order")
	}

	images, err := repositories.GetImagesByTimestamp(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if len(images) == 0 {
		return c.JSON(http.StatusNotFound, "No images found")
	}

	if order < 0 || order >= len(images) {
		return c.JSON(http.StatusBadRequest, "Invalid image order")
	}

	selectedImage := images[order]

	c.Response().Header().Set("Content-Type", "image/png")
	c.Response().Write(selectedImage.Mading)
	fmt.Println("Gambar Berhasil Dicomot")
	return nil
}

