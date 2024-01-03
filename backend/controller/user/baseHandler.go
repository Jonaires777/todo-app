package user

import (
	"todo-app/config"
	"todo-app/services/user"
)

type BaseHandler struct {
	Service *user.Service
	Validator *config.RequestBodyValidator 
}

func NewBaseHandler(
	service *user.Service,
	validator *config.RequestBodyValidator,
) *BaseHandler {
	return &BaseHandler{
		Service: service,
		Validator: validator,
	}
}