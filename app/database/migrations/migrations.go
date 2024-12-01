package migrations

import (
	"database/sql"

	"github.com/lmd1e/song_library/app/utils"
)

func RunMigrations(db *sql.DB) error {
	utils.Logger.Info("Running database migrations")
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS songs (
            id SERIAL PRIMARY KEY,
            "group" VARCHAR(255) NOT NULL,
            song VARCHAR(255) NOT NULL,
            release_date TIMESTAMP NOT NULL,
            text TEXT NOT NULL,
            link VARCHAR(255) NOT NULL
        );
    `)
	if err != nil {
		utils.Logger.Error("Failed to run migrations: ", err)
		return err
	}

	utils.Logger.Info("Migrations completed successfully")
	return nil
}
