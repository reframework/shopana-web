package httpServer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewServer() *fiber.App {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		// TODO: allow based on env
		AllowOrigins: "*",
	}))

	return app
}
