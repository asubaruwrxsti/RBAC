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
	// Check if the token is valid
	if !parsedToken.Valid {
		fmt.Println("Token is invalid")
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	userId := parsedToken.Claims.(jwt.MapClaims)["userId"]
	groupId := int(parsedToken.Claims.(jwt.MapClaims)["groupId"].(float64))
	fmt.Println("userId: ", userId)
	fmt.Println("groupId: ", groupId)

	// Simulate where the user does not have permission to access this resource
	if groupId == 2 {
		fmt.Println("User does not have permission to access this resource, groupId: ", groupId)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if err != nil {
		fmt.Printf("Error parsing token: %v\n", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	fmt.Println("<< Token is valid, returning ...")

	return c.JSON(fiber.Map{
		"userId":  userId,
		"groupId": groupId,
		"token":   token,
		"message": "Token is valid",
	})
}
