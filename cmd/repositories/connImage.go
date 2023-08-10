package repositories

import (
	"fmt"
	"golang/cmd/storage"
	"golang/cmd/models"

	_ "database/sql"
	_ "github.com/lib/pq"
)

func GetImage() ([]models.Image, error) {
	db := storage.GetDB()

	sqlStatement := `
		SELECT id, name FROM mading
	`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve images: %w", err)
	}
	defer rows.Close()

	var images []models.Image

	for rows.Next() {
		var image models.Image
		if err := rows.Scan(&image.Id, &image.Name); err != nil {
			return nil, fmt.Errorf("failed to scan image: %w", err)
		}
		images = append(images, image)
	}

	return images, nil
}

func UploadImage(name string, imageData []byte) error {
	db := storage.GetDB()

	sqlStatement := `INSERT INTO mading (name, mading) VALUES ($1, $2)`

	_, err := db.Exec(sqlStatement, name, imageData)
	if err != nil {
		return fmt.Errorf("failed to upload image: %w", err)
	}

	return nil
}