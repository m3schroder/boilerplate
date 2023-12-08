package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"strukt/database"
	_ "strukt/docs"
	"strukt/middleware"
	"strukt/router"
)

// @title           App API Docs
// @version         1.0
// @description     Base setup for swagger

// @contact.name   Matthew Schroder
// @contact.email  m3schroder@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "App Name",
	})

	// Database Connections
	database.ConnectPostgres("DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME")

	// Middleware
	app.Use(recover.New())
	app.Use(middleware.Logger())
	app.Get("/swagger/*", middleware.Swagger())
	app.Get("/monitor", middleware.Monitor())

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
