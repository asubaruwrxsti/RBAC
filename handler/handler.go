package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func ValidateToken(c *fiber.Ctx) error {
	fmt.Println(">> Inside ValidateToken")
	// Read the value of the "Authorization" header
	token := c.Get("Authorization")[7:]

	// Parse the JWT token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		fmt.Printf("Error parsing token: %v\n", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"userId":  parsedToken.Claims.(jwt.MapClaims)["userId"],
		"groupId": parsedToken.Claims.(jwt.MapClaims)["groupId"],
		"token":   token,
		"message": "Token is valid",
	})
}
