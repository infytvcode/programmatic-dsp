package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Tracking(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
	})
}
