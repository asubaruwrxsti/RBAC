package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func ParseToken(c *fiber.Ctx) error {
	// Read the value of the "Authorization" header, slice the bearer
	token := c.Get("Authorization")[7:]

	// Parse the JWT token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	// Check if the token is valid
	if !parsedToken.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	userId := parsedToken.Claims.(jwt.MapClaims)["userId"]
	groupId := int(parsedToken.Claims.(jwt.MapClaims)["groupId"].(float64))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"userId":  userId,
		"groupId": groupId,
		"token":   token,
		"message": "Token is valid",
	})
}
