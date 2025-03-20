package contact

import (
	"log"
	"myapp/internal/DTO"

	"github.com/gofiber/fiber/v2"
)

func GetContact(c *fiber.Ctx) error {
	log.Printf("Get contact")
	contact := DTO.ContactResponse{
		ID:         1,
		Username:   "Test",
		GivenName:  "GiventName",
		FamilyName: "FamilyName",
	}
	return c.Status(fiber.StatusOK).JSON(contact)
}

func AddContact(c *fiber.Ctx) error {
	log.Printf("Add contact")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

func UpdateContact(c *fiber.Ctx) error {
	log.Printf("Update contact")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

func DeleteContact(c *fiber.Ctx) error {
	log.Printf("Delete contact")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}
