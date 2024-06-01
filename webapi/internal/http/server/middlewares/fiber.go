package middlewares

import (
	ctxutil "webapi/internal/app/ctx"
	app "webapi/internal/app/interface"

	"github.com/gofiber/fiber/v2"
)

func FiberMiddleWare(_ app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Locals(ctxutil.FiberCtxKey) != nil {
			return c.Next()
		}

		c.Locals(ctxutil.FiberCtxKey, c)
		return c.Next()
	}
}
