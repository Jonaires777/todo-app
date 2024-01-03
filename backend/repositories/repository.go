package repositories

import (
	"todo-app/repositories/todo"
	"todo-app/repositories/user"

	"gorm.io/gorm"
)

type Repository = struct {
	DB *gorm.DB
	User user.Repository
	Todo todo.Repository
}

func SetupRepository(connection *gorm.DB) *Repository{
	return &Repository{
		DB: connection,
		User: user.NewUserRepository(connection),
		Todo: todo.NewTodoRepository(connection),
	}
}