package middlewares

import (
	ctxutil "webapi/internal/app/ctx"
	app "webapi/internal/app/interface"
	appErrors "webapi/pkg/error"
	"webapi/pkg/helpers"

	"github.com/gofiber/fiber/v2"
)

func TenantMiddleware(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Locals(ctxutil.UserKey) != nil {
			return c.Next()
		}

		jwt := helpers.GetBearerToken(c.Get("Authorization"))
		if jwt == "" {
			return c.Next()
		}

		user, err := app.Service().User().Me(ctxutil.FromFiber(c), jwt)
		if err != nil {
			return c.Status(401).JSON(appErrors.Unauthorized.New(err))
		}

		c.Locals(ctxutil.UserKey, user)
		return c.Next()
	}
}
