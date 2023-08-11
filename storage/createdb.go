package storage

import (
	"fmt"
)

func CreateDB() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS tipex (
			tipeid SERIAL PRIMARY KEY,
			nama VARCHAR NOT NULL
		);
	`)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Table berhasil dibuat.")
}
