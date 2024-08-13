package todo

import (
	"strconv"
	"time"
)

type Todo struct {
	ID      int
	Text    string
	Created time.Time
}

func (t Todo) Slice() []string {
	return []string{strconv.Itoa(t.ID), t.Text, t.Created.Format(time.RFC3339)}
}
