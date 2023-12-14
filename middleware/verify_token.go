package middleware

import (
	"RBAC/config"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			// No token, return error
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Parse the token
		// Remove Bearer from token (first 7 characters)
		if len(tokenString) > len("Bearer ") && tokenString[:len("Bearer ")] == "Bearer " {
			// Remove the "Bearer " prefix
			tokenString = tokenString[len("Bearer "):]
		} else {
			// If the prefix is not present, handle it accordingly
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Applying the signing method check
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return []byte(config.Config("JWT_SECRET")), nil
		})

		switch {
		case token.Valid:
			return c.Next()
		case errors.Is(err, jwt.ErrTokenMalformed):
			// Token is malformed
			return c.SendStatus(fiber.StatusUnauthorized)
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			// Invalid signature
			return c.SendStatus(fiber.StatusUnauthorized)
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			// Token is either expired or not active yet
			return c.SendStatus(fiber.StatusUnauthorized)
		default:
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}
}
