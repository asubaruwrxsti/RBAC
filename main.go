package main

import (
	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2" // import the fiber package
	"github.com/gofiber/template/html/v2"

	"log"

	"RBAC/database"
	"RBAC/router"

	_ "github.com/lib/pq"

	"github.com/gofiber/swagger"
	"go.uber.org/zap"
)

// @title Go Fiber REST API with Swagger Example
// @description Go Fiber REST API with Swagger Example
// @version 1
// @host localhost:3000
// @BasePath /
// @schemes http
func main() {

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	engine := html.New("./views", ".html")
	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)

	logger, _ := zap.NewProduction()

	// app.Use(middleware.Logger())
	router.SetupRoutes(app)

	app.Get("/swagger/*", swagger.HandlerDefault)     // default
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://localhost:3000/swagger/doc.json",
		DeepLinking: false,
	}))

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))

	log.Fatal(app.Listen(":3000"))

}
