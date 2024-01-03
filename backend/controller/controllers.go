package controller

import (
	"todo-app/config"
	"todo-app/controller/user"
	"todo-app/services"
)

type Controllers struct {
	User user.BaseHandler
}

func SetupControllers(service *services.Service) *Controllers {
	validator := config.NewValidator()

	return &Controllers{
		User: *user.NewBaseHandler(&service.User, validator),
	}
}