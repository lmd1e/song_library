package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lmd1e/song_library/app/controllers"
	"github.com/lmd1e/song_library/app/models"
	"github.com/lmd1e/song_library/app/requests"
	"github.com/lmd1e/song_library/app/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetSongs(t *testing.T) {
	mockRepo := new(mocks.MockSongRepository)
	songController := controllers.NewSongController(mockRepo)

	expectedSongs := []models.Song{
		{ID: 1, Group: "Test Group", Song: "Test Song", ReleaseDate: time.Now(), Text: "Test song text", Link: "https://example.com/test-song"},
	}
	mockRepo.On("GetSongs", mock.AnythingOfType("map[string]string"), 10, 0).Return(expectedSongs, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/songs?group=Test%20Group", nil)

	router := gin.Default()
	router.GET("/songs", songController.GetSongs)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var songs []models.Song
	json.Unmarshal(w.Body.Bytes(), &songs)
	assert.Equal(t, 1, len(songs))
	assert.Equal(t, "Test Group", songs[0].Group)
	assert.Equal(t, "Test Song", songs[0].Song)

	mockRepo.AssertExpectations(t)
}

func TestGetSongText(t *testing.T) {

	mockRepo := new(mocks.MockSongRepository)
	songController := controllers.NewSongController(mockRepo)

	expectedText := "Test song text"
	mockRepo.On("GetSongText", 1, 10, 0).Return(expectedText, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/songs/1/text", nil)

	router := gin.Default()
	router.GET("/songs/:id/text", songController.GetSongText)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, expectedText, response["text"])

	mockRepo.AssertExpectations(t)
}

func TestDeleteSong(t *testing.T) {

	mockRepo := new(mocks.MockSongRepository)
	songController := controllers.NewSongController(mockRepo)

	mockRepo.On("DeleteSong", 1).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/songs/1", nil)

	router := gin.Default()
	router.DELETE("/songs/:id", songController.DeleteSong)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Song deleted", response["message"])

	mockRepo.AssertExpectations(t)
}

func TestUpdateSong(t *testing.T) {

	mockRepo := new(mocks.MockSongRepository)
	songController := controllers.NewSongController(mockRepo)

	updatedSong := models.Song{
		ID:          1,
		Group:       "Updated Group",
		Song:        "Updated Song",
		ReleaseDate: time.Now(),
		Text:        "Updated song text",
		Link:        "https://example.com/updated-song",
	}
	mockRepo.On("UpdateSong", mock.MatchedBy(func(song models.Song) bool {
		return song.ID == updatedSong.ID &&
			song.Group == updatedSong.Group &&
			song.Song == updatedSong.Song &&
			song.Text == updatedSong.Text &&
			song.Link == updatedSong.Link
	})).Return(nil)

	w := httptest.NewRecorder()
	jsonBody, _ := json.Marshal(updatedSong)
	req, _ := http.NewRequest("PUT", "/songs/1", bytes.NewBuffer(jsonBody))

	router := gin.Default()
	router.PUT("/songs/:id", songController.UpdateSong)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var responseSong models.Song
	json.Unmarshal(w.Body.Bytes(), &responseSong)
	assert.Equal(t, updatedSong.Group, responseSong.Group)
	assert.Equal(t, updatedSong.Song, responseSong.Song)

	mockRepo.AssertExpectations(t)
}

func TestAddSong(t *testing.T) {

	mockRepo := new(mocks.MockSongRepository)
	songController := controllers.NewSongController(mockRepo)

	newSong := models.Song{
		Group:       "New Group",
		Song:        "New Song",
		ReleaseDate: time.Now(),
		Text:        "New song text",
		Link:        "https://example.com/new-song",
	}
	mockRepo.On("AddSong", mock.MatchedBy(func(song models.Song) bool {
		return song.Group == newSong.Group &&
			song.Song == newSong.Song &&
			song.Text == newSong.Text &&
			song.Link == newSong.Link
	})).Return(nil)

	w := httptest.NewRecorder()
	songRequest := requests.AddSongRequest{
		Group: "New Group",
		Song:  "New Song",
	}
	jsonBody, _ := json.Marshal(songRequest)
	req, _ := http.NewRequest("POST", "/songs", bytes.NewBuffer(jsonBody))

	router := gin.Default()
	router.POST("/songs", songController.AddSong)

	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	var responseSong models.Song
	json.Unmarshal(w.Body.Bytes(), &responseSong)
	assert.Equal(t, newSong.Group, responseSong.Group)
	assert.Equal(t, newSong.Song, responseSong.Song)

	mockRepo.AssertExpectations(t)
}
