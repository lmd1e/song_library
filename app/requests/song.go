package requests

import (
	"encoding/json"
	"net/http"
	"os"
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
	externalAPIURL := os.Getenv("EXTERNAL_API_URL")
	resp, err := http.Get(externalAPIURL + group + "&song=" + song)
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
