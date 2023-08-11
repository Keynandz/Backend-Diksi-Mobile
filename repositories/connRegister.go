package repositories

import (
	"fmt"
	"golang/models"
	"golang/storage"
)

func CreateAkun(akun models.Akun) (models.Akun, error) {
	db := storage.GetDB()
	sqlStatement := `INSERT INTO akun (username, password, email, phone) VALUES ($1, $2, $3, $4) RETURNING id`

	err := db.QueryRow(sqlStatement, akun.Username, akun.Password, akun.Email, akun.Phone).Scan(&akun.Id)
	if err != nil {
		return akun, fmt.Errorf("error creating akun: %w", err)
	}

	return akun, nil
}
