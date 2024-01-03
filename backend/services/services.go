package services

import (
	"todo-app/repositories"
	"todo-app/services/todo"
	"todo-app/services/user"
)

type Service struct {
	User user.Service
	Todo todo.Service
}

func SetupeServices(repository *repositories.Repository) *Service {
	return &Service{
		User: *user.SetupService(&repository.User),
		Todo: *todo.SetupTodoService(&repository.Todo),
	}
}