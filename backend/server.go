package main

import (
	"log"
	"todo-app/config"
	"todo-app/controller"
	"todo-app/services"

	"github.com/gofiber/fiber/v2"
)

func main(){	
	server := fiber.New()

	repository, err := config.SetupDB(".env")

	if err != nil {
		log.Fatal(err)
	}

	defer config.CloseDB(repository)

	services := services.SetupeServices(repository)

	controllers := controller.SetupControllers(services)

	server.Post("/signup", controllers.User.CreateUser)

	server.Listen(":8000")
}

