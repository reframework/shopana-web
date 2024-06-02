package graphqlHandler

import (
	"context"
	"errors"

	// app "portal/services/project/internal/app/interface"
	// "portal/services/project/internal/delivery/http/graphql-core/resolvers"
	// "portal/services/project/internal/delivery/http/server/middlewares"

	app "webapi/internal/app/interface"
	appErrors "webapi/pkg/error"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/kr/pretty"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// todo: pass an app
func InitRouter(router *fiber.App, app app.App) {
	srv := handler.NewDefaultServer(nil /* graph_gen.NewExecutableSchema(config) */)

	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)

		appErr, ok := errors.Unwrap(err.Err).(*appErrors.AppError)
		if ok {
			err.Message = appErr.Message
			err.Extensions = map[string]any{
				"code":         appErr.Code,
				"descriptions": appErr.Descriptions,
				"caller":       appErr.Caller,
				"version":      appErr.Version,
			}
		}

		return err
	})

	srv.AddTransport(transport.POST{})
	srv.Use(extension.Introspection{})

	pretty.Log("Admin API: POST /api/admin/graphql/query")

	api := router.Group("/api/admin/graphql")
	// api.Use(middlewares.LoaderMiddleware(app))
	api.Post("/query", adaptor.HTTPHandler(srv))

	pretty.Log("Admin playground: GET /api/admin/graphql")

	router.Get("/api/admin/graphql", adaptor.HTTPHandlerFunc(
		playground.Handler("Pixli admin GraphQL", "/api/admin/graphql/query")))
}
