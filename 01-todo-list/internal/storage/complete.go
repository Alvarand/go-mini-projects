package storage

import (
	"errors"
)

var errorTaskIsAlreadyCompleted = errors.New("task is already completed")
var errorTaskIsNotExists = errors.New("taskID is not exists")

func (s *storage) Complete(taskID int) error {
	isExistTaskID := false

	todos := s.List()
	for i, todo := range todos {
		if todo.ID != taskID || todo.IsDeleted {
			continue
		}
		if todo.IsCompleted {
			return errorTaskIsAlreadyCompleted
		}
		todos[i].IsCompleted = true
		isExistTaskID = true
		break
	}

	if !isExistTaskID {
		return errorTaskIsNotExists
	}

	return s.rewrite(todos)
}
