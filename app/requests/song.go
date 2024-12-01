package requests

import (
	"encoding/json"
	"net/http"
	"time"
)

type AddSongRequest struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type SongDetail struct {
	ReleaseDate time.Time `json:"releaseDate"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
}

func GetSongDetail(group, song string) (*SongDetail, error) {
	resp, err := http.Get("http://external-api.com/info?group=" + group + "&song=" + song)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var songDetail SongDetail
	if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
		return nil, err
	}

	return &songDetail, nil
}
