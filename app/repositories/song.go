package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/lmd1e/song_library/models"
)

type SongRepository struct {
	db *sql.DB
}

func NewSongRepository(db *sql.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) GetSongs(filter map[string]string, limit, offset int) ([]models.Song, error) {
	query := "SELECT id, group, song, release_date, text, link FROM songs"
	var conditions []string
	var args []interface{}
	i := 1
	for k, v := range filter {
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
		return nil, err
	}
	defer rows.Close()

	var songs []models.Song
	for rows.Next() {
		var song models.Song
		if err := rows.Scan(&song.ID, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link); err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}

func (r *SongRepository) GetSongText(songID, limit, offset int) (string, error) {
	query := "SELECT text FROM songs WHERE id = $1"
	var text string
	err := r.db.QueryRow(query, songID).Scan(&text)
	if err != nil {
		return "", err
	}
	// Пагинация текста песни
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

func (r *SongRepository) DeleteSong(songID int) error {
	query := "DELETE FROM songs WHERE id = $1"
	_, err := r.db.Exec(query, songID)
	return err
}

func (r *SongRepository) UpdateSong(song models.Song) error {
	query := `
        UPDATE songs
        SET group = $1, song = $2, release_date = $3, text = $4, link = $5
        WHERE id = $6
    `
	_, err := r.db.Exec(query, song.Group, song.Song, song.ReleaseDate, song.Text, song.Link, song.ID)
	return err
}

func (r *SongRepository) AddSong(song models.Song) error {
	query := `
        INSERT INTO songs (group, song, release_date, text, link)
        VALUES ($1, $2, $3, $4, $5)
    `
	_, err := r.db.Exec(query, song.Group, song.Song, song.ReleaseDate, song.Text, song.Link)
	return err
}
