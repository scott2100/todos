package todo

import (
	"strconv"
	"time"
)

type Todo struct {
	ID          int
	Description string
	Created     time.Time
	Completed   time.Time
}

func (t Todo) Slice() []string {
	return []string{strconv.Itoa(t.ID), t.Description, t.Created.Format(time.RFC3339), t.Completed.Format(time.RFC3339)}
}
