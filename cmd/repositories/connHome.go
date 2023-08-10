package repositories

import (
	"database/sql"
	"fmt"
	"golang/cmd/models"
	"golang/cmd/storage"
)


func GetAkunByID(identifier string) (models.Akun, error) {
	db := storage.GetDB()

	sqlStatement := `
		SELECT id, username, email, password FROM akun WHERE username = $1 OR email = $1
`

	row := db.QueryRow(sqlStatement, identifier)

	user := models.Akun{}
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}
		return user, fmt.Errorf("failed to retrieve user: %w", err)
	}

	return user, nil
}
