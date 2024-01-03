package todo

import (
	"time"
	"todo-app/models"

	"gorm.io/gorm"
)

type (
	Repository struct {
		GetBD func() *gorm.DB
	}

	ICreate struct {
		Title string
		Status uint
		CreatedBy uint
		Deadline time.Time
	}

	IUpdate struct {
		ID uint
		Title *string
		CreatedBy *uint
		Status *uint
		Deadline *time.Time
	}

	IDelete struct {
		ID uint
	}

	Interface interface {
		Create(ICreate) (*models.Todos, error)
		ReadAll() ([]models.Todos, error)
		Update(IUpdate) (*models.Todos, error)
		Delete(IDelete) (*models.Todos, error)
	}
)