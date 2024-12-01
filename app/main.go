package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lmd1e/song_library/controllers"
	"github.com/lmd1e/song_library/database"
	"github.com/lmd1e/song_library/repositories"
	"github.com/lmd1e/song_library/routes"

	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := database.RunMigrations(db); err != nil {
		log.Fatal(err)
	}

	songRepo := repositories.NewSongRepository(db)

	songController := controllers.NewSongController(songRepo)

	router := mux.NewRouter()
	routes.RegisterSongRoutes(router, songController)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
