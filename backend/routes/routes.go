package routes

import (
	"todo-app/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", handlers.GetTodos)
	app.Get("/todo/:id", handlers.GetTodo)
	app.Post("/create", handlers.NewTodo)
	app.Delete("/todo/:id", handlers.DeleteTodo)
	
}