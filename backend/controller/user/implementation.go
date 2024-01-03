package user

import "github.com/gofiber/fiber/v2"

type CreateUserResquestBody struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (handler *BaseHandler) CreateUser(c *fiber.Ctx) error {
	body := new(CreateUserResquestBody)

	if parseError := c.BodyParser(body); parseError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request body",
			"error": parseError.Error(),
		})
	}

	validationErrors := handler.Validator.Validate(body)

	if validationErrors != nil {
		errorMessages := make([]string, len(validationErrors))

		for index := range validationErrors {
			errorMessages[index] = validationErrors[index].Message
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Invalid request body",
			"errors": errorMessages,
		})
	}

	newUser, serviceError := handler.Service.NewUser(
		body.Name,
		body.Email,
		body.Password,
	)

	if serviceError != nil  {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user creation error",
			"error": serviceError.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created successfully",
		"user": newUser,
	})
}
