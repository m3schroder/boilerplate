package router

import (
	"github.com/gofiber/fiber/v2"

	"strukt/handler"
	"strukt/middleware"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {

	// Base
	v1 := app.Group("/api/v1", handler.Hello)
	{
		// Auth
		auth := v1.Group("/auth")
		{
			auth.Post("/login", handler.Login)
		}

		// User
		user := v1.Group("/user")
		{
			user.Get("/:id", handler.GetUser)
			user.Post("/", handler.CreateUser)
			user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
			user.Delete("/:id", middleware.Protected(), handler.DeleteUser)
		}
	}

}
