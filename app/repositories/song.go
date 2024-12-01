package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/lmd1e/song_library/app/models"
	"github.com/lmd1e/song_library/app/utils"
)

type SongRepository interface {
	GetSongs(filter map[string]string, limit, offset int) ([]models.Song, error)
	GetSongText(songID, limit, offset int) (string, error)
	DeleteSong(songID int) error
	UpdateSong(song models.Song) error
	AddSong(song models.Song) error
}

type SongRepositoryImpl struct {
	db *sql.DB
}

func NewSongRepository(db *sql.DB) *SongRepositoryImpl {
	return &SongRepositoryImpl{db: db}
}

func (r *SongRepositoryImpl) GetSongs(filter map[string]string, limit, offset int) ([]models.Song, error) {
	utils.Logger.Info("Fetching songs from the database")
	query := "SELECT id, \"group\", song, release_date, text, link FROM songs"
	var conditions []string
	var args []interface{}
	i := 1
	for k, v := range filter {
		if k == "group" {
			k = `"group"`
		}
		conditions = append(conditions, fmt.Sprintf("%s = $%d", k, i))
		args = append(args, v)
		i++
	}
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)
	rows, err := r.db.Query(query, args...)
	if err != nil {
		utils.Logger.Error("Failed to fetch songs: ", err)
		return nil, err
	}
	defer rows.Close()

	var songs []models.Song
	for rows.Next() {
		var song models.Song
		if err := rows.Scan(&song.ID, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link); err != nil {
			utils.Logger.Error("Failed to scan song row: ", err)
			return nil, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}

func (r *SongRepositoryImpl) GetSongText(songID, limit, offset int) (string, error) {
	utils.Logger.Info("Fetching song text from the database")
	query := "SELECT text FROM songs WHERE id = $1"
	var text string
	err := r.db.QueryRow(query, songID).Scan(&text)
	if err != nil {
		utils.Logger.Error("Failed to fetch song text: ", err)
		return "", err
	}

	lines := strings.Split(text, "\n")
	if offset >= len(lines) {
		return "", nil
	}
	end := offset + limit
	if end > len(lines) {
		end = len(lines)
	}
	return strings.Join(lines[offset:end], "\n"), nil
}

func (r *SongRepositoryImpl) DeleteSong(songID int) error {
	utils.Logger.Info("Deleting song from the database")
	query := "DELETE FROM songs WHERE id = $1"
	_, err := r.db.Exec(query, songID)
	if err != nil {
		utils.Logger.Error("Failed to delete song: ", err)
	}
	return err
}

func (r *SongRepositoryImpl) UpdateSong(song models.Song) error {
	utils.Logger.Info("Updating song in the database")
	query := `
        UPDATE songs
        SET "group" = $1, song = $2, release_date = $3, text = $4, link = $5
        WHERE id = $6
    `
	_, err := r.db.Exec(query, song.Group, song.Song, song.ReleaseDate, song.Text, song.Link, song.ID)
	if err != nil {
		utils.Logger.Error("Failed to update song: ", err)
	}
	return err
}

func (r *SongRepositoryImpl) AddSong(song models.Song) error {
	utils.Logger.Info("Adding song to the database")
	query := `
        INSERT INTO songs ("group", song, release_date, text, link)
        VALUES ($1, $2, $3, $4, $5)
    `
	_, err := r.db.Exec(query, song.Group, song.Song, song.ReleaseDate, song.Text, song.Link)
	if err != nil {
		utils.Logger.Error("Failed to add song: ", err)
	}
	return err
}
