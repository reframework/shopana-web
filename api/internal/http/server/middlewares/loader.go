package middlewares

import (
	app "webapi/internal/app/interface"
	"webapi/internal/loader"

	"github.com/gofiber/fiber/v2"
)

func LoaderMiddleware(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals(loader.CONTEXT_KEY, loader.New(app))

		return c.Next()
	}
}
