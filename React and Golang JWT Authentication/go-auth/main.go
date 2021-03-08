package main

import (
	"github.com/gofiber/fiber/v2"
	"go-auth/databese"
	"go-auth/routes"
)

func main() {
	databese.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
