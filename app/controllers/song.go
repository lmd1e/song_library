package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lmd1e/song_library/app/docs"
	"github.com/lmd1e/song_library/app/models"
	"github.com/lmd1e/song_library/app/repositories"
	"github.com/lmd1e/song_library/app/requests"
	"github.com/lmd1e/song_library/app/utils"
)

type SongController struct {
	repo repositories.SongRepository
}

func NewSongController(repo repositories.SongRepository) *SongController {
	return &SongController{repo: repo}
}

// @Summary Получение данных библиотеки с фильтрацией и пагинацией
// @Description Получение данных библиотеки с фильтрацией по всем полям и пагинацией
// @Tags Songs
// @Accept json
// @Produce json
// @Param group query string false "Фильтр по группе"
// @Param song query string false "Фильтр по названию песни"
// @Param limit query int false "Количество записей на странице"
// @Param offset query int false "Смещение (страница)"
// @Success 200 {array} models.Song
// @Failure 500 {object} map[string]string
// @Router /songs [get]
func (c *SongController) GetSongs(ctx *gin.Context) {
	utils.Logger.Info("GetSongs request received")
	filter := make(map[string]string)
	for key, value := range ctx.Request.URL.Query() {
		filter[key] = value[0]
	}
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	songs, err := c.repo.GetSongs(filter, limit, offset)
	if err != nil {
		utils.Logger.Error("Failed to fetch songs: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch songs"})
		return
	}
	ctx.JSON(http.StatusOK, songs)
}

// @Summary Получение текста песни с пагинацией по куплетам
// @Description Получение текста песни с пагинацией по куплетам
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Param limit query int false "Количество куплетов на странице"
// @Param offset query int false "Смещение (страница)"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /songs/{id}/text [get]
func (c *SongController) GetSongText(ctx *gin.Context) {
	utils.Logger.Info("GetSongText request received")
	songID, _ := strconv.Atoi(ctx.Param("id"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	text, err := c.repo.GetSongText(songID, limit, offset)
	if err != nil {
		utils.Logger.Error("Failed to fetch song text: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch song text"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"text": text})
}

// @Summary Удаление песни
// @Description Удаление песни по ID
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /songs/{id} [delete]
func (c *SongController) DeleteSong(ctx *gin.Context) {
	utils.Logger.Info("DeleteSong request received")
	songID, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.repo.DeleteSong(songID); err != nil {
		utils.Logger.Error("Failed to delete song: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete song"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Song deleted"})
}

// @Summary Изменение данных песни
// @Description Изменение данных песни по ID
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Param song body models.Song true "Данные песни"
// @Success 200 {object} models.Song
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /songs/{id} [put]
func (c *SongController) UpdateSong(ctx *gin.Context) {
	utils.Logger.Info("UpdateSong request received")
	songID, _ := strconv.Atoi(ctx.Param("id"))
	var song models.Song
	if err := ctx.ShouldBindJSON(&song); err != nil {
		utils.Logger.Error("Invalid request payload: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	song.ID = songID
	if err := c.repo.UpdateSong(song); err != nil {
		utils.Logger.Error("Failed to update song: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update song"})
		return
	}
	ctx.JSON(http.StatusOK, song)
}

// @Summary Добавление новой песни
// @Description Добавление новой песни в формате JSON
// @Tags Songs
// @Accept json
// @Produce json
// @Param song body requests.AddSongRequest true "Данные песни"
// @Success 201 {object} models.Song
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /songs [post]
func (c *SongController) AddSong(ctx *gin.Context) {
	utils.Logger.Info("AddSong request received")
	var req requests.AddSongRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Logger.Error("Invalid request payload: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	songDetail, err := requests.GetSongDetail(req.Group, req.Song)
	if err != nil {
		utils.Logger.Error("Failed to fetch song details: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch song details"})
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
		utils.Logger.Error("Failed to add song: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add song"})
		return
	}

	ctx.JSON(http.StatusCreated, song)
}
