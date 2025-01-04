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
		log.Fatalf(models.ErrorFailedOpenFile.Error(), err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			log.Fatalf(models.ErrorFailedCloseFile.Error(), err)
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
		log.Fatalf(models.ErrorFailedWriteFile.Error(), err)
	}

	return id
}
