package models

import (
	"fmt"
	"time"
)

type TODO struct {
	ID          int
	IsCompleted bool
	IsDeleted   bool
	Description string
	CreatedAt   time.Time
}

func (t TODO) GetValues() []string {
	return []string{fmt.Sprint(t.ID), fmt.Sprint(t.IsCompleted), fmt.Sprint(t.IsDeleted), fmt.Sprint(t.Description), t.CreatedAt.Format(TimeLayout)}
}
