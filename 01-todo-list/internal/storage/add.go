package storage

import (
	"encoding/csv"
	"log"
	"os"
	"time"
	"todo-list/internal/models"
)

func (s *storage) Add(description string) int {
	fi, err := os.OpenFile(s.path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			log.Fatalf("failed to close file: %s", err)
		}
	}()

	id := s.getMaxID() + 1
	todo := models.TODO{
		ID:          id,
		Description: description,
		CreatedAt:   time.Now(),
	}

	csvWriter := csv.NewWriter(fi)
	defer csvWriter.Flush()

	err = csvWriter.Write(todo.GetValues())
	if err != nil {
		log.Fatalf("failed to write in file: %s", err)
	}

	return id
}
