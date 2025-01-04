package storage

import (
	"errors"
	"log"
	"os"
)

const path = "./input.csv"

type Storage interface {
	Add(string) int
	Delete()
	List()
	Complete()
}

type storage struct {
	path string
}

func (s *storage) getMaxID() int {
	// TODO генерировать инкремент ID
	return 0
}

var s *storage

func init() {
	s = &storage{
		path: path,
	}

	if _, err := os.Stat(s.path); errors.Is(err, os.ErrNotExist) {
		fo, err := os.Create(s.path)
		if err != nil {
			log.Fatalf("failed to create file: %s", err)
		}
		defer func() {
			if err := fo.Close(); err != nil {
				log.Fatalf("failed to close file: %s", err)
			}
		}()
	}
}

func GetStorage() Storage {
	return s
}
