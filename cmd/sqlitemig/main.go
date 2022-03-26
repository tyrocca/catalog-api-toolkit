package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/tyrocca/catalog-api-toolkit/config"
)

func main() {
	db, _, err := config.InitializeDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Beginning migrations")
	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fSrc, err := (&file.File{}).Open("./db/migrations/sqlite")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to sql")

	m, err := migrate.NewWithInstance("file", fSrc, "sqlite3", instance)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Beginning UP migration")
	// modify for Down
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
	log.Println("Finished Migration")
}
