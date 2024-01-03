package todo

import (
	"todo-app/models"
	"todo-app/repositories/todo"
)

type (
	Service struct {
		TodoRepo *todo.Repository
	}

	Interface interface {
		MarkTodoAsCompleted(id uint) (*models.Todos, error)
	}
)