package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lmd1e/song_library/app/controllers"
	migrations "github.com/lmd1e/song_library/app/database/migrations"
	"github.com/lmd1e/song_library/app/repositories"
	"github.com/lmd1e/song_library/app/routes"
	"github.com/lmd1e/song_library/app/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/lib/pq"
)

// @title Song Library API
// @version 1.0
// @description API для онлайн библиотеки песен
// @host localhost:8080
// @BasePath /
func main() {

	if err := godotenv.Load(); err != nil {
		utils.Logger.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		utils.Logger.Fatal(err)
	}
	defer db.Close()

	if err := migrations.RunMigrations(db); err != nil {
		utils.Logger.Fatal(err)
	}

	songRepo := repositories.NewSongRepository(db)

	songController := controllers.NewSongController(songRepo)

	router := gin.Default()
	routes.RegisterSongRoutes(router, songController)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	utils.Logger.Info("Server started on :8080")
	log.Fatal(router.Run(":8080"))
}
