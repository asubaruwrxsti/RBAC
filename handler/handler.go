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

	// parse the jwt token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		fmt.Printf("Error parsing token: %v\n", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Print the token to the console (optional)
	fmt.Printf("Token from header: %s\n", token)
	// get the claims
	fmt.Printf("Claims -> UserID: %+v\n", parsedToken.Claims.(jwt.MapClaims)["userId"])
	fmt.Printf("Claims -> GroupId: %+v\n", parsedToken.Claims.(jwt.MapClaims)["groupId"])
	return c.JSON(fiber.Map{
		"userId":  parsedToken.Claims.(jwt.MapClaims)["userId"],
		"groupId": parsedToken.Claims.(jwt.MapClaims)["groupId"],
		"token":   token,
		"message": "Token is valid",
	})
}
