package todo

import (
	"todo-app/config"
	"todo-app/services/todo"
)

type BaseHandler struct {
	Service *todo.Service
	Validator *config.RequestBodyValidator
}

func NewBaseHandler(service *todo.Service, validator *config.RequestBodyValidator) *BaseHandler{
	return &BaseHandler{
		Service: service,
		Validator: validator,
	}
}