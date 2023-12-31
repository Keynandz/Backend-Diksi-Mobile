package handlers

import (
	"fmt"
	"golang/session"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LogoutAkun(c echo.Context) error {
	session.ClearSession(c)
	fmt.Println("clearing session...")
	return c.JSON(http.StatusOK, "Logout successful")
}