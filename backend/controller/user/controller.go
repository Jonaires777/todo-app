package user

import (
	userService "todo-app/services/user"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	UserService userService.Interface
}

func New(userService userService.Interface) *Controller {
	return &Controller{
		UserService: userService,
	}
}

func (uc *Controller) Create(c *fiber.Ctx) error {
	var createUser userService.ICreateUser
	if err := c.BodyParser(&createUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	newUser, err := uc.UserService.NewUser(createUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(newUser)
}
