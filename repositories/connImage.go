package repositories

import (
	"fmt"
	"golang/models"
	"golang/storage"
	"time"

	_ "database/sql"

	_ "github.com/lib/pq"
)

func UploadImage(name string, mading []byte) error {
	db := storage.GetDB()

	sqlStatement := `INSERT INTO mading (name, mading, update_at) VALUES ($1, $2, $3)`

	_, err := db.Exec(sqlStatement, name, mading, time.Now())
	if err != nil {
		return fmt.Errorf("failed to upload image: %w", err)
	}

	return nil
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
