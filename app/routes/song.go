package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lmd1e/song_library/app/controllers"
)

func RegisterSongRoutes(router *gin.Engine, controller *controllers.SongController) {
	router.GET("/songs", controller.GetSongs)
	router.GET("/songs/:id/text", controller.GetSongText)
	router.DELETE("/songs/:id", controller.DeleteSong)
	router.PUT("/songs/:id", controller.UpdateSong)
	router.POST("/songs", controller.AddSong)
}
