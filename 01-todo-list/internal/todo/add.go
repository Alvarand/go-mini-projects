package todo

import (
	"log"
	"todo-list/internal/models"
	"todo-list/internal/storage"
)

func Add(args []string) {
	models.CheckZeroArgs(args)
	description := args[0]
	taskID := storage.GetStorage().Add(description)

	log.Println("Created new todo task with ID:", taskID)
}
