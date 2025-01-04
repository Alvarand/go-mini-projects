package storage

func (s *storage) Delete(taskID int) error {
	isExistTaskID := false

	todos := s.List()
	for i, todo := range todos {
		if todo.ID != taskID || todo.IsDeleted {
			continue
		}
		todos[i].IsDeleted = true
		isExistTaskID = true
		break
	}

	if !isExistTaskID {
		return errorTaskIsNotExists
	}

	return s.rewrite(todos)
}
