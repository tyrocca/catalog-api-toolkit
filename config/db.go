package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeDB() (*sql.DB, bool, error) {
	// Initialize connection to the database
	initialized := false
	db, err := sql.Open("sqlite3", "./catalogapi.db")

	if err != nil {
		log.Println("Driver creation failed", err.Error())
		return db, initialized, err
	} else {
		log.Println("Database driver created!")
	}

	_, select_error := db.Query("SELECT 1 FROM catalog_item LIMIT 1")
	if select_error != nil {
		log.Println("DB not migrated", select_error.Error())
	}

	return db, initialized, err
}
