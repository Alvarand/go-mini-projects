package storage

import (
	"encoding/csv"
	"log"
	"os"
	"todo-list/internal/models"
)

func (s *storage) rewrite(todos []models.TODO) error {
	fi, err := os.OpenFile(s.path, os.O_TRUNC|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalf(models.ErrorFailedOpenFile.Error(), err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			log.Fatalf(models.ErrorFailedCloseFile.Error(), err)
		}
	}()

	csvWriter := csv.NewWriter(fi)
	defer csvWriter.Flush()

	for _, todo := range todos {
		err = csvWriter.Write(todo.GetValues())
		if err != nil {
			log.Fatalf(models.ErrorFailedWriteFile.Error(), err)
		}
	}

	return nil
}
