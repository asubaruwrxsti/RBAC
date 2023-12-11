package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"Title":   "Hello, World!",
		"Message": "Hello, World!",
	})
}
