package repository

import (
	"database/sql"
	"todo-app/repositories/user"
	"todo-app/repositories/todo"
)

type Repository struct {
	DB *sql.DB
	User user.Repository
	Todo todo.Repository
}

func SetupRepository(connection *sql.DB) *Repository{
	return &Repository{
		DB:	connection,
		User: user.NewUserRepository(connection),
		Todo: todo.NewTodoRepository(connection),
	}
}