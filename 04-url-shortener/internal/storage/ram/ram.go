package ram

import (
	"errors"
	"sync"
)

const countOfAlphabet = 52

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
	shortURL := ld.generateURL()
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

func (ld *LocalDatabase) generateURL() string {
	num := ld.maxID
	reverseGeneratedURL := ""

	for num > 0 {
		alphabetIndex := num % countOfAlphabet
		if alphabetIndex == 0 {
			alphabetIndex = countOfAlphabet
			num = num / (countOfAlphabet + 1)
		} else {
			num /= countOfAlphabet
		}

		if alphabetIndex > countOfAlphabet/2 {
			reverseGeneratedURL += string(rune('A' - 1 + alphabetIndex - (countOfAlphabet / 2)))
		} else {
			reverseGeneratedURL += string(rune('a' - 1 + alphabetIndex))
		}
	}

	shortURL := ""
	for i := len(reverseGeneratedURL) - 1; 0 <= i; i-- {
		shortURL += string(reverseGeneratedURL[i])
	}
	ld.maxID++
	return shortURL
}
