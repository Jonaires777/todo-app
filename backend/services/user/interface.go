package user

import (
	"todo-app/models"
	"todo-app/repositories/user"
)

const (
	emailMaxLen = 128
	emailMinLen = 3
	nameMaxLen  = 128
	nameMinLen  = 3

	passwordMinLen = 8
)

type (
	Service struct {
		UserRepo *user.Repository
	}

	ICreateUser struct {
		name     string
		email    string
		password string
	}

	Interface interface {
		NewUser(ICreateUser) (*models.Users, error)
	}
)
