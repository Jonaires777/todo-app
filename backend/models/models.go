package models

import "time"

const (
	TStatusFinished = 1
	TStatusPending  = 2
	TStatusOverdue  = 3
)

var TodoStatus = map[uint]string{
	TStatusFinished: "finished",
	TStatusPending:  "pending",
	TStatusOverdue:  "overdue",
}

type Users struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Todos struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Status    uint      `json:"status"`
	CreatedBy uint      `json:"createdby"`
	Deadline  time.Time `json:"deadline"`
}
