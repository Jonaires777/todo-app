package handlers

import (
	"todo-app/models"
	"todo-app/repositories"

	"github.com/gofiber/fiber/v2"
)

func GetTodos(c *fiber.Ctx) error {	
	db := repositories.DBConn
	var todos []models.Todos
	db.Find(&todos)
	return c.JSON(todos)
}

func GetTodo(c *fiber.Ctx) error{
	id := c.Params("id")
	db := repositories.DBConn
	var todo models.Todos
	db.Find(&todo, id)
	return c.JSON(todo)
}

func NewTodo(c *fiber.Ctx) error {
	db := repositories.DBConn
	todo := new(models.Todos)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&todo)
	return c.JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := repositories.DBConn

	var todo models.Todos
	db.First(&todo, id)
	if todo.Title == "" {
		return c.Status(500).SendString("No Todo found With Id")
	}
	db.Delete(&todo)
	return c.SendString("Book deleted successfully")
}

func NewUser(c *fiber.Ctx) error {
	db := repositories.DBConn
	user := new(models.Users)
	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&user)
	return c.JSON(user)
}

