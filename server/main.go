package main

import (
	"new_go_app/src/database"
	"new_go_app/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	routes.Routes(app)

	app.Listen(":3000")
}
