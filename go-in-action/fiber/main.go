package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/poneding/learning-go/practice/fiber/config"

	"github.com/poneding/learning-go/practice/fiber/handlers"
)

func main() {
	app := fiber.New()

	config.Connect()

	app.Get("/users", handlers.GetUsers)
	app.Get("/users/:id", handlers.GetUser)
	app.Post("/users", handlers.AddUser)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.RemoveUser)

	log.Fatal(app.Listen(":3000"))
}
