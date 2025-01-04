package todo

import (
	"errors"
	"log"
	"strconv"
	"todo-list/internal/models"
	"todo-list/internal/storage"
)

var errorFailedDelete = errors.New("failed to delete taskID: %s")

func Delete(args []string) {
	models.CheckZeroArgs(args)
	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf(models.ErrorConvertStringToInt.Error(), err)
	}

	err = storage.GetStorage().Delete(taskID)
	if err != nil {
		log.Fatalf(errorFailedDelete.Error(), err)
	}

	log.Printf("Successful delete taskID = %d\n", taskID)
}
