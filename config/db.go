package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitializeDB() (*sql.DB, error) {
	// Initialize connection to the database
	var err error
	DB, err = sql.Open("sqlite3", "./catalogapi.db")

	if err != nil {
		log.Println("Driver creation failed", err.Error())
	} else {
		log.Println("Database driver created!")
	}

	_, err = DB.Query("SELECT * FROM userinfo")
	if err != nil {
		log.Println("Driver creation failed", err.Error())
	} else {
		log.Println("errrrrr")
	}

	return DB, err
}
