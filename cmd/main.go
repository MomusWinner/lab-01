package main

import (
	"myapp/internal/api/v1/contact"
	"myapp/internal/api/v1/group"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	api := app.Group("/api/v1")

	api.Route("/", contact.AddRoutes)
	api.Route("/", group.AddRoutes)
	app.Listen(":7080")
}
