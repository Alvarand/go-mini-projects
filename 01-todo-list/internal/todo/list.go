package todo

import (
	"log"
	"todo-list/internal/storage"
)

func List(_ []string) {
	todos := storage.GetStorage().List()
	for _, todo := range todos {
		if todo.IsDeleted {
			continue
		}
		log.Println(todo)
	}
}
