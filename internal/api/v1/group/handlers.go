package group

import (
	"log"
	"myapp/internal/DTO"

	"github.com/gofiber/fiber/v2"
)

func GetGroup(c *fiber.Ctx) error {
	contact := DTO.GroupResponse{}
	log.Printf("Get group")
	return c.Status(fiber.StatusOK).JSON(contact)
}

func AddGroup(c *fiber.Ctx) error {
	log.Printf("Add group")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

func UpdateGroup(c *fiber.Ctx) error {
	log.Printf("Update group")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

func DeleteGroup(c *fiber.Ctx) error {
	log.Printf("Delete group")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}
