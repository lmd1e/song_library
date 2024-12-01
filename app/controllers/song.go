package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lmd1e/song_library/models"
	"github.com/lmd1e/song_library/repositories"
	"github.com/lmd1e/song_library/requests"
	"github.com/lmd1e/song_library/utils"
)

type SongController struct {
	repo *repositories.SongRepository
}

func NewSongController(repo *repositories.SongRepository) *SongController {
	return &SongController{repo: repo}
}

func (c *SongController) GetSongs(w http.ResponseWriter, r *http.Request) {
	filter := make(map[string]string)
	for key, value := range r.URL.Query() {
		filter[key] = value[0]
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	songs, err := c.repo.GetSongs(filter, limit, offset)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch songs")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, songs)
}

func (c *SongController) GetSongText(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	songID, _ := strconv.Atoi(vars["id"])
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	text, err := c.repo.GetSongText(songID, limit, offset)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch song text")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"text": text})
}

func (c *SongController) DeleteSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	songID, _ := strconv.Atoi(vars["id"])
	if err := c.repo.DeleteSong(songID); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to delete song")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Song deleted"})
}

func (c *SongController) UpdateSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	songID, _ := strconv.Atoi(vars["id"])
	var song models.Song
	if err := utils.DecodeJSONBody(w, r, &song); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	song.ID = songID
	if err := c.repo.UpdateSong(song); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update song")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, song)
}

func (c *SongController) AddSong(w http.ResponseWriter, r *http.Request) {
	var req requests.AddSongRequest
	if err := utils.DecodeJSONBody(w, r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	songDetail, err := requests.GetSongDetail(req.Group, req.Song)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch song details")
		return
	}

	song := models.Song{
		Group:       req.Group,
		Song:        req.Song,
		ReleaseDate: songDetail.ReleaseDate,
		Text:        songDetail.Text,
		Link:        songDetail.Link,
	}

	if err := c.repo.AddSong(song); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to add song")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, song)
}
