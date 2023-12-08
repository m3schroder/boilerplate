package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Swagger() fiber.Handler {
	return swagger.HandlerDefault
}
