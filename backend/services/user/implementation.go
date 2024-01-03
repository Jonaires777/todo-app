package user

import (
	"todo-app/models"
	"todo-app/repositories/user"
)

func SetupService(repositoryUser *user.Repository) *Service {
	return &Service{
		UserRepo: repositoryUser,
	}
}

func (service *Service) NewUser(name, email, password string) (*models.Users, error) {
	hashedPassword, err := HashPassword(password)

	if err != nil {
		return nil, err
	}

	newUser, err := service.UserRepo.Create(user.ICreate{
		Name: name,
		Email: email,
		Password: hashedPassword,
	})

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

