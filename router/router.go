package router

import (
	"RBAC/handler"
	"RBAC/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, dbConn *gorm.DB) {
	authMiddleware := middleware.NewAuth().
		Use(middleware.RequestToken()).
		Use(middleware.VerifyToken())

	app.Get("/", handler.Home)
	app.Get("parse", handler.ParseToken)

	app.Get("/authreq", authMiddleware.Apply, func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
}
