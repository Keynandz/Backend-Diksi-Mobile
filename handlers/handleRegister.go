package handlers

import (
	"golang/models"
	"golang/repositories"
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateAkun(c echo.Context) error {
    akun := models.Akun{}
    c.Bind(&akun)

    // Validate the phone number using a regular expression
    if !isValidPhoneNumber(akun.Phone) {
        return c.JSON(http.StatusBadRequest, "Invalid phone number. Only numeric characters are allowed.")
    }

    if err := encryptPassword(&akun); err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    newAkun, err := repositories.CreateAkun(akun)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusCreated, newAkun)
}

func isValidPhoneNumber(phone string) bool {
    // Define a regular expression to match only numeric characters
    numericRegex := regexp.MustCompile(`^[0-9]+$`)
    return numericRegex.MatchString(phone)
}


func encryptPassword(akun *models.Akun) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(akun.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	akun.Password = string(hashedPassword)
	return nil
}
