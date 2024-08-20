package todo

import (
	"strconv"
	"time"
)

type Todo struct {
	ID          int
	Description string
	Created     time.Time
	IsComplete  bool
}

func (t Todo) Slice() []string {
	return []string{strconv.Itoa(t.ID), t.Description, t.Created.Format(time.RFC3339), strconv.FormatBool(t.IsComplete)}
}
