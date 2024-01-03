package user

import (
	"todo-app/models"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) Repository {
	return Repository{
		GetBD: func() *gorm.DB {
			return conn
		},
	}
}

func (r *Repository) Create(createData ICreate) (*models.Users, error) {
	var user = models.Users{
		Name: createData.Name,
		Email: createData.Email,
		Password: createData.Password,
	}

	result := r.GetBD().Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}