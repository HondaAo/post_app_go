package main

import (
	"new_go_app/server/src/database"
	"new_go_app/server/src/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Routes(app)
	os.Setenv("ENV", "Development")

	app.Listen(":4000")
}
