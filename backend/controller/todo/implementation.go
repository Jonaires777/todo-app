package todo

import "github.com/gofiber/fiber/v2"

type CreateUserResquestBodyId struct {
	Id uint `json:"id" validate:"required"`
}

func (handler *BaseHandler) MarkTodoAsCompleted(c *fiber.Ctx) error {
	body := new(CreateUserResquestBodyId)

	if parseError := c.BodyParser(body); parseError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Invalid request body",
			"error": parseError.Error(),
		})
	}

	validateErrors := handler.Validator.Validate(body)

	if validateErrors != nil {
		errorsMessage := make([]string, len(validateErrors))

		for index := range validateErrors {
			errorsMessage[index] = validateErrors[index].Message
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error": errorsMessage,
		})
	}

	newTodo, err := handler.Service.MarkTodoAsCompleted(body.Id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Todo updating error",
			"error": err.Error(),
		})	
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Todo marked as complete sucessfully",
		"todo": newTodo,
	})
}

