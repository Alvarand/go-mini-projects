package todo

import (
	"log"
	"strconv"
	"todo-list/internal/models"
	"todo-list/internal/storage"
)

func Delete(args []string) {
	models.CheckZeroArgs(args)
	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("failed to convert string to int: %s", err)
	}

	err = storage.GetStorage().Delete(taskID)
	if err != nil {
		log.Fatalf("failed to delete taskID: %s", err)
	}

	log.Printf("Successful delete taskID = %d\n", taskID)
}
