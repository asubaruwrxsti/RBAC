package main

import (
	"github.com/gofiber/fiber/v2" // import the fiber package

	"log"

	"RBAC/database"
	"RBAC/router"
)

func main() {

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{})

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
