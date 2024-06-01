package ctxutil

import (
	"context"

	app "webapi/internal/app/interface"
	"webapi/internal/entity"

	"github.com/gofiber/fiber/v2"
)

type Key string

const (
	RequestIdKey Key = "request_id"
	UserKey      Key = "current_user"
	LoadersKey   Key = "loaders"
	FiberCtxKey  Key = "fiber_ctx"
)

var keys = []Key{
	RequestIdKey,
	UserKey,
	LoadersKey,
	FiberCtxKey,
}

func FromFiber(fiberCtx *fiber.Ctx) context.Context {
	ctx := context.Background()

	for _, key := range keys {
		ctx = context.WithValue(ctx, key, fiberCtx.Locals(key))
	}

	return ctx
}

func Fiber(ctx context.Context, fiberCtx *fiber.Ctx) context.Context {
	for _, key := range keys {
		ctx = context.WithValue(ctx, key, fiberCtx.Locals(string(key)))
	}

	return ctx
}

func User(ctx context.Context) (*entity.User, bool) {
	user, ok := ctx.Value(UserKey).(*entity.User)
	return user, ok
}

func SetUser(ctx context.Context, customer *entity.User) context.Context {
	return context.WithValue(ctx, UserKey, customer)
}

func FiberCtx(ctx context.Context) (*fiber.Ctx, bool) {
	fiberCtx, ok := ctx.Value(FiberCtxKey).(*fiber.Ctx)
	return fiberCtx, ok
}

func RequestId(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(RequestIdKey).(string)
	return user, ok
}

func SetRequestId(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, UserKey, id)
}

func Loader(ctx context.Context) (app.Loader, bool) {
	loader, ok := ctx.Value(LoadersKey).(app.Loader)
	return loader, ok
}

func SetLoader(ctx context.Context, loader app.Loader) context.Context {
	return context.WithValue(ctx, LoadersKey, loader)
}
