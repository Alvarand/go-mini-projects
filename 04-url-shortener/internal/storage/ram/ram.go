package ram

import (
	"errors"
	"sync"
)

var errorUnexistsShortURL = errors.New("url does not exist")

type LocalDatabase struct {
	maxID   int
	storage map[string]string
	mx      sync.Mutex
}

func New() *LocalDatabase {
	storage := make(map[string]string)
	return &LocalDatabase{
		maxID:   1,
		storage: storage,
		mx:      sync.Mutex{},
	}
}

func (ld *LocalDatabase) SaveURL(url string) (string, error) {
	ld.mx.Lock()
	shortURL := "" // TODO generate short URL
	ld.storage[shortURL] = url
	ld.mx.Unlock()

	return shortURL, nil
}

func (ld *LocalDatabase) GetURL(shortURL string) (string, error) {
	ld.mx.Lock()
	url, isExistURL := ld.storage[shortURL]
	ld.mx.Unlock()

	if !isExistURL {
		return "", errorUnexistsShortURL
	}

	return url, nil
}
