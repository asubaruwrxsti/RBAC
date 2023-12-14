package middleware

import (
	"RBAC/config"
	"RBAC/database"
	"RBAC/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RequestToken() func(*fiber.Ctx) error {
	dbConn := database.DB

	return func(c *fiber.Ctx) error {
		// Later to be verified from the DB !!!
		username := config.Config("AUTH_USERNAME")
		password := config.Config("AUTH_PASSWORD")
		// Check if the username and password are correct
		if c.FormValue("username") == username && c.FormValue("password") == password {

			// Get the user id
			user := model.User{}
			row := dbConn.Table("users").Select("id").Where("full_name = ?", username).Row()
			err := row.Scan(&user.ID)
			if err != nil {
				return c.SendStatus(fiber.StatusUnauthorized)
			}

			// Get the user group
			userGroup := model.Group{}
			row = dbConn.Table("user_groups").Select("id").Where("user_id = ?", user.ID).Row()
			err = row.Scan(&userGroup.ID)
			if err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			}

			claims := jwt.MapClaims{
				"userId":  user.ID,       // Custom claims
				"groupId": userGroup.ID,  // Custom claims
				"iss":     "your_issuer", // Set the issuer
				"exp":     time.Now().Add(time.Hour * 72).Unix(),
			}

			// Create a new token object, specifying signing method and the claims
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			// Sign and get the complete encoded token as a string using the secret
			tokenString, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
			if err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			}
			return c.JSON(fiber.Map{"token": tokenString})
		}
		return c.SendStatus(fiber.StatusUnauthorized)
	}
}
