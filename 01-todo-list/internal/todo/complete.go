package todo

import (
	"log"
	"strconv"
	"todo-list/internal/models"
	"todo-list/internal/storage"
)

func Complete(args []string) {
	models.CheckZeroArgs(args)
	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("failed to convert string to int: %s", err)
	}

	err = storage.GetStorage().Complete(taskID)
	if err != nil {
		log.Fatalf("failed to complete taskID: %s", err)
	}

	log.Printf("Successful complete taskID = %d\n", taskID)
}
