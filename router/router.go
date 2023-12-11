package router

import (
	"RBAC/handler"
	"RBAC/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// Public routes
	app.Get("/", handler.Home)
	app.Get("/auth", middleware.Auth())
}
