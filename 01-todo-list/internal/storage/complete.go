package storage

import (
	"errors"
)

func (s *storage) Complete(taskID int) error {
	isExistTaskID := false

	todos := s.List()
	for i, todo := range todos {
		if todo.ID != taskID || todo.IsDeleted {
			continue
		}
		if todo.IsCompleted {
			return errors.New("task is already completed")
		}
		todos[i].IsCompleted = true
		isExistTaskID = true
		break
	}

	if !isExistTaskID {
		return errors.New("taskID is not exists")
	}

	return s.rewrite(todos)
}
