package middlewares

import (
	ctxutil "webapi/internal/app/ctx"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ReqIdMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals(ctxutil.RequestIdKey, uuid.New().String())
		return c.Next()
	}
}
