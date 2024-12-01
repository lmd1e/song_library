package mocks

import (
	"github.com/lmd1e/song_library/app/models"
	"github.com/stretchr/testify/mock"
)

type MockSongRepository struct {
	mock.Mock
}

func (m *MockSongRepository) GetSongs(filter map[string]string, limit, offset int) ([]models.Song, error) {
	args := m.Called(filter, limit, offset)
	return args.Get(0).([]models.Song), args.Error(1)
}

func (m *MockSongRepository) GetSongText(songID, limit, offset int) (string, error) {
	args := m.Called(songID, limit, offset)
	return args.String(0), args.Error(1)
}

func (m *MockSongRepository) DeleteSong(songID int) error {
	args := m.Called(songID)
	return args.Error(0)
}

func (m *MockSongRepository) UpdateSong(song models.Song) error {
	args := m.Called(song)
	return args.Error(0)
}

func (m *MockSongRepository) AddSong(song models.Song) error {
	args := m.Called(song)
	return args.Error(0)
}
