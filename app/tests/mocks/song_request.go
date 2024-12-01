package mocks

import (
	"github.com/lmd1e/song_library/app/requests"
	"github.com/stretchr/testify/mock"
)

type MockSongRequest struct {
	mock.Mock
}

func (m *MockSongRequest) GetSongDetail(group, song string) (*requests.SongDetail, error) {
	args := m.Called(group, song)
	return args.Get(0).(*requests.SongDetail), args.Error(1)
}
