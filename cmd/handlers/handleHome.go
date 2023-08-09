package handlers

import (
	"net/http"
	"go-collab/cmd/repositories"

	"github.com/labstack/echo/v4"
)

func GetAkunByID(c echo.Context) error {
	akunIdentifier := c.Param("identifier") // Assuming you get the identifier from the route
	existingAkun, err := repositories.GetAkunByID(akunIdentifier)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, existingAkun)
}