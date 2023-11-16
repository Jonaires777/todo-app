package models

import "time"

const (
	TStatusFinished = 1
	TStatusPending  = 2
	TStatusOverdue = 3
)

var TodoStatus = map[uint]string{
	TStatusFinished: "finished",
	TStatusPending:  "pending",
	TStatusOverdue: "overdue",
}

type Todos struct {
	ID      int
	Title   string
	Status  uint
	Content string
	Created time.Time
}
