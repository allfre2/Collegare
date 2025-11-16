package services

import (
	"collegare/domain/models"
	"io"
	"os"
	"sync"
)

type MockFileCache struct {
	mutex    sync.RWMutex
	cache    map[string][]byte
	filesDir string
}

func NewMockFileCache() *MockFileCache {
	return &MockFileCache{
		cache:    make(map[string][]byte),
		filesDir: "./poc/audio",
	}
}

func (c *MockFileCache) GetTrack(track models.Track) ([]byte, error) {
	if b, ok := c.cache[track.Filename]; ok {
		return b, nil
	}

	filename := c.filesDir + "/" + track.Filename

	_, err := os.Stat(filename)
	if !os.IsNotExist(err) {
		return nil, err
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache[track.Filename] = data

	return data, nil
}
