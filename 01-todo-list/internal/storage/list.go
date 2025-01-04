package storage

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
	"todo-list/internal/models"
)

func (s *storage) List() []models.TODO {

	todos := make([]models.TODO, 0)

	fo, err := os.OpenFile(s.path, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		log.Fatalf("failed to open csv file: %s", err)
	}

	defer func() {
		if err := fo.Close(); err != nil {
			log.Fatalf("failed to close csv file: %s", err)
		}
	}()

	csvReader := csv.NewReader(fo)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("failed to read csv file: %s", err)
	}

	for _, record := range records {
		todos = append(todos, parseTODO(record))
	}

	return todos
}

func parseTODO(record []string) models.TODO {
	id := stringToInt(record[0])
	isCompleted := stringToBool(record[1])
	isDeleted := stringToBool(record[2])
	description := record[3]
	createdAt := stringToTime(record[4])

	return models.TODO{
		ID:          id,
		IsCompleted: isCompleted,
		IsDeleted:   isDeleted,
		Description: description,
		CreatedAt:   createdAt,
	}
}

func stringToInt(s string) int {
	parsedInt, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to convert string to int: %s", err)
	}
	return parsedInt
}

func stringToBool(s string) bool {
	parsedBool, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatalf("failed to convert string to bool: %s", err)
	}
	return parsedBool
}

func stringToTime(s string) time.Time {
	parsedTime, err := time.Parse(models.TimeLayout, s)
	if err != nil {
		log.Fatalf("failed to convert string to time: %s", err)
	}

	return parsedTime
}
