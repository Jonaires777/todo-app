package models

import "time"

type Users struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}