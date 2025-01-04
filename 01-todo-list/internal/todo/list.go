package todo

import (
	"log"
	"todo-list/internal/storage"
)

func List(_ []string) {
	todos := storage.GetStorage().List()
	for _, todo := range todos {
		log.Println(todo)
	}
}
