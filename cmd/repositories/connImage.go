package repositories

import (
	"fmt"
	"golang/cmd/models"
	"golang/cmd/storage"
	"time"

	_ "database/sql"

	_ "github.com/lib/pq"
)

func GetImage() ([]models.Image, error) {
	db := storage.GetDB()

	sqlStatement := `
		SELECT imageid, name FROM mading
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

func UploadImage(name string, mading []byte) error {
	db := storage.GetDB()

	sqlStatement := `INSERT INTO mading (name, mading, update_at) VALUES ($1, $2, $3)`

	_, err := db.Exec(sqlStatement, name, mading, time.Now())
	if err != nil {
		return fmt.Errorf("failed to upload image: %w", err)
	}

	return nil
}

func GetImageByID(id string) (*models.Image, error) {
	db := storage.GetDB()

	sqlStatement := `
        SELECT imageid, name, mading FROM mading WHERE imageid = $1
    `

	row := db.QueryRow(sqlStatement, id)

	image := &models.Image{}
	err := row.Scan(&image.Id, &image.Name, &image.Mading)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve image: %w", err)
	}

	return image, nil
}

func GetImagesByTimestamp(order int) ([]models.Image, error) {
    db := storage.GetDB()

    sqlStatement := `
        SELECT imageid, name, mading FROM mading ORDER BY update_at DESC
    `

    rows, err := db.Query(sqlStatement)
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve images: %w", err)
    }
    defer rows.Close()

    var images []models.Image

    for rows.Next() {
        var image models.Image
        if err := rows.Scan(&image.Id, &image.Name, &image.Mading); err != nil {
            return nil, fmt.Errorf("failed to scan image: %w", err)
        }
        images = append(images, image)
    }

    return images, nil
}
