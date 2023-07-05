package main

import (
	"log"

	"fiber-sample/config"
	"fiber-sample/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	err := config.Connect()
	if err != nil {
		log.Fatal(err)
	}

	app.Get("/users", handlers.GetUsers)
	app.Get("/users/:id", handlers.GetUser)
	app.Post("/users", handlers.AddUser)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.RemoveUser)

	log.Fatal(app.Listen(":3000"))
}
