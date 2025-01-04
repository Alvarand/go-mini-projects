package storage

import (
	"errors"
	"log"
	"os"
	"todo-list/internal/models"
)

const path = "./input.csv"

type Storage interface {
	Add(string) int
	Delete(int) error
	List() []models.TODO
	Complete(int) error
}

type storage struct {
	path string
}

func (s *storage) getMaxID() (maxID int) {
	todos := s.List()

	for _, todo := range todos {
		if todo.ID > maxID {
			maxID = todo.ID
		}
	}
	return maxID
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
