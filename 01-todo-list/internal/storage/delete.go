package storage

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
)

func (s *storage) Delete(taskID int) error {
	isExistTaskID := false

	todos := s.List()
	for i, todo := range todos {
		if todo.ID != taskID || todo.IsDeleted {
			continue
		}
		todos[i].IsDeleted = true
		isExistTaskID = true
		break
	}

	if !isExistTaskID {
		return errors.New("taskID is not exists")
	}

	fi, err := os.OpenFile(s.path, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalf("failed to open csv file: %s", err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			log.Fatalf("failed to close csv file: %s", err)
		}
	}()

	csvWriter := csv.NewWriter(fi)
	defer csvWriter.Flush()

	for _, todo := range todos {
		err = csvWriter.Write(todo.GetValues())
		if err != nil {
			log.Fatalf("failed to write in file: %s", err)
		}
	}

	return nil
}
