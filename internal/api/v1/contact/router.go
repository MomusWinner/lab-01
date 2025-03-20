package contact

import "github.com/gofiber/fiber/v2"

func AddRoutes(router fiber.Router) {
	contact := router.Group("contact")
	contact.Post("/", AddContact)
	contact.Get("/", GetContact)
	contact.Delete("/", DeleteContact)
	contact.Put("/", UpdateContact)
}
