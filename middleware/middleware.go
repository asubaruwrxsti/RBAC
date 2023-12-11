package middleware

import (
	"RBAC/config"
	"RBAC/database"
	"RBAC/model"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func RequestToken() func(*fiber.Ctx) error {
	dbConn := database.DB

	return func(c *fiber.Ctx) error {
		username := config.Config("AUTH_USERNAME")
		password := config.Config("AUTH_PASSWORD")
		log.Print(">> Inside Auth middleware")

		// Check if the username and password are correct
		if c.FormValue("username") == username && c.FormValue("password") == password {

			// Get the user id
			user := model.User{}
			row := dbConn.Table("users").Select("id").Where("full_name = ?", username).Row()
			err := row.Scan(&user.ID)
			if err != nil {
				log.Print("<< Auth middleware: ", err)
				return c.SendStatus(fiber.StatusUnauthorized)
			}

			if user.ID == 0 {
				log.Print("<< Auth middleware: userId is 0")
				return c.SendStatus(fiber.StatusUnauthorized)
			}

			// Get the user group
			userGroup := model.Group{}
			row = dbConn.Table("user_groups").Select("id").Where("user_id = ?", user.ID).Row()
			err = row.Scan(&userGroup.ID)
			if err != nil {
				log.Print("<< Auth middleware: ", err)
				return c.SendStatus(fiber.StatusInternalServerError)
			}

			if userGroup.ID == 0 {
				log.Print("<< Auth middleware: userGroup is empty")
				return c.SendStatus(fiber.StatusUnauthorized)
			}

			// Create a new token object, specifying signing method and the claims
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"userId":  user.ID,
				"groupId": userGroup,
				"iss":     "your_issuer", // Set the issuer
				"exp":     time.Now().Add(time.Hour * 72).Unix(),
			})

			// Sign and get the complete encoded token as a string using the secret
			tokenString, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
			if err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			}
			log.Print("<< Auth middleware tokenString: ", tokenString)
			return c.JSON(fiber.Map{"token": tokenString})
		}

		log.Print(">> Auth middleware: username or password is incorrect")
		return c.SendStatus(fiber.StatusUnauthorized)
	}
}

func VerifyToken() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		log.Print(">> Inside AuthReq middleware")
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			// No token, return error
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Remove Bearer from token
		// Parse the token
		tokenString = tokenString[7:]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Make sure that the token method conforms to "SigningMethodHMAC"
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return []byte(config.Config("JWT_SECRET")), nil
		})
		log.Printf("<< AuthReq middleware token: %+v\n", token)
		log.Printf("<< AuthReq middleware err: %+v\n", err)

		if err != nil {
			// Handle specific errors, e.g., token expired, issuer mismatch, etc.
			switch err.(type) {
			case *jwt.ValidationError:
				return c.SendStatus(fiber.StatusUnauthorized)
			default:
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		log.Print("<< AuthReq middleware token.Valid: ", token.Valid)
		if !token.Valid {
			// Invalid token, return error
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Token is valid, proceed with the request
		return c.Next()
	}
}
