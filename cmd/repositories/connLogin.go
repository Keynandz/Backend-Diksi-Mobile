package repositories

import (
	"fmt"
	"go-collab/cmd/models"
	"go-collab/cmd/storage"
)

func GetUserByEmailOrUsername(identifier string) (*models.Akun, error) {
	db := storage.GetDB()
	akun := &models.Akun{}
	sqlStatement := `SELECT id, username, password, email, phone FROM akun WHERE email = $1 OR username = $2`

	err := db.QueryRow(sqlStatement, identifier, identifier).Scan(&akun.Id, &akun.Username, &akun.Password, &akun.Email, &akun.Phone)
	if err != nil {
		return nil, fmt.Errorf("error fetching akun: %w", err)
	}

	return akun, nil
}
