package user

import (
	"todo-app/models"

	"gorm.io/gorm"
)

const (
	emailMaxLenght    = 128
	emailMinLenght    = 3
	nameMaxLenght     = 128
	nameMinLenght     = 3
	passwordMinLength = 8
)

type (
	Repository struct {
		GetBD func() *gorm.DB
	}

	ICreate struct{
		Name string
		Email string
		Password string
	}

	Interface interface{
		Create(ICreate) (*models.Users, error)
	}
)