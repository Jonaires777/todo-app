package todo

import (
	"todo-app/models"
	"todo-app/repositories/todo"
)

func SetupTodoService(repository *todo.Repository) *Service {
	return &Service{
		TodoRepo: repository,
	}
}

func (service *Service) MarkTodoAsCompleted(id uint) (*models.Todos, error) {
	status := uint(1)
	statusp := &status

	todoUpdate, err := service.TodoRepo.Update(todo.IUpdate{
		ID: id,
		Status: statusp,
	})
	
	if err != nil {
		return nil, err
	}

	return todoUpdate, nil
}

