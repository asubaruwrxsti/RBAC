package router

import (
	"RBAC/handler"
	"RBAC/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.Home)
	app.Get("/auth", middleware.RequestToken())
	app.Get("/authreq", middleware.VerifyToken(), handler.ValidateToken)
}
