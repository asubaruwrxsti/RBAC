package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func ValidateToken(c *fiber.Ctx) error {
	fmt.Println(">> Inside ValidateToken")
	return c.SendString("Hello 👋! Your token is valid!")
}
