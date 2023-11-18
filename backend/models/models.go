package models

import (
	"gorm.io/gorm"
)

type Users struct {
	ID       uint
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserModel struct {
	DB *gorm.DB
}

type Todos struct {
	ID     uint
	Title  string `json:"title"`
	Status uint   `json:"status"`
}

type TodoModel struct {
	DB *gorm.DB
}
