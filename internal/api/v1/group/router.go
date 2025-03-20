package group

import "github.com/gofiber/fiber/v2"

func AddRoutes(router fiber.Router) {
	contact := router.Group("group")
	contact.Post("/", AddGroup)
	contact.Get("/", GetGroup)
	contact.Delete("/", DeleteGroup)
	contact.Put("/", UpdateGroup)
}
