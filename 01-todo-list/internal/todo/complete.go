package todo

import (
	"errors"
	"log"
	"strconv"
	"todo-list/internal/models"
	"todo-list/internal/storage"
)

var errorFailedComplete = errors.New("failed to complete taskID: %s")

func Complete(args []string) {
	models.CheckZeroArgs(args)
	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf(models.ErrorConvertStringToInt.Error(), err)
	}

	err = storage.GetStorage().Complete(taskID)
	if err != nil {
		log.Fatalf(errorFailedComplete.Error(), err)
	}

	log.Printf("Successful complete taskID = %d\n", taskID)
}
