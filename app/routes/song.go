package routes

import (
	"github.com/gorilla/mux"
	"github.com/lmd1e/song_library/controllers"
)

func RegisterSongRoutes(router *mux.Router, controller *controllers.SongController) {
	router.HandleFunc("/songs", controller.GetSongs).Methods("GET")
	router.HandleFunc("/songs/{id}/text", controller.GetSongText).Methods("GET")
	router.HandleFunc("/songs/{id}", controller.DeleteSong).Methods("DELETE")
	router.HandleFunc("/songs/{id}", controller.UpdateSong).Methods("PUT")
	router.HandleFunc("/songs", controller.AddSong).Methods("POST")
}
