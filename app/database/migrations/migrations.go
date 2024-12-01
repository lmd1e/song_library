package database

import (
	"database/sql"
	"log"
)

func RunMigrations(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS songs (
            id SERIAL PRIMARY KEY,
            group VARCHAR(255) NOT NULL,
            song VARCHAR(255) NOT NULL,
            release_date TIMESTAMP NOT NULL,
            text TEXT NOT NULL,
            link VARCHAR(255) NOT NULL
        );
    `)
	if err != nil {
		return err
	}

	log.Println("Migrations completed successfully")
	return nil
}
