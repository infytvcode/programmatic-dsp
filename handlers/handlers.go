package handlers

import (
	"github.com/infytvcode/dsp/database"
	"github.com/infytvcode/dsp/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// UserList returns a list of users
func CustomerList(c *fiber.Ctx) error {
	customers := database.Get()

	return c.JSON(fiber.Map{
		"success":   true,
		"customers": customers,
	})
}

// UserCreate registers a user
func CustomerCreate(c *fiber.Ctx) error {
	customer := &models.Customer{
		// Note: when writing to external database,
		// we can simply use - Name: c.FormValue("user")
		Name: utils.CopyString(c.FormValue("user")),
	}
	database.Insert(customer)

	return c.JSON(fiber.Map{
		"success": true,
		"user":    customer,
	})
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
