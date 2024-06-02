package loader

import (
	"context"
	"fmt"
	"strings"

	app "webapi/internal/app/interface"

	"github.com/graph-gophers/dataloader"
)

type (
	LoaderContextKey string
	LoaderEnum       string
	LoadersMap       map[LoaderEnum]*dataloader.Loader
)

const (
	CONTEXT_KEY LoaderContextKey = "dataLoaderV3"
	USER        LoaderEnum       = "user"
)

type Loader struct {
	Loaders LoadersMap
}

func New(app app.App) *Loader {
	return &Loader{
		Loaders: LoadersMap{},
	}
}

func Load[T any](ctx context.Context, loaderKey LoaderEnum, key string) (T, error) {
	loader, ok := ctx.Value(CONTEXT_KEY).(*Loader)

	var z T
	if !ok {
		return z, fmt.Errorf("loader: not found")
	}

	l, ok := loader.Loaders[loaderKey]
	if !ok {
		return z, fmt.Errorf("loader: loader %s does not exist", loaderKey)
	}

	thunk := l.Load(ctx, dataloader.StringKey(key))
	result, err := thunk()
	if err != nil {
		return z, err
	}

	value, ok := result.(T)
	if !ok {
		return z, fmt.Errorf("could not convert result to T. Loader key: %s", key)
	}

	return value, nil
}

func MakeKey[T comparable](k T, fields ...string) string {
	if len(fields) == 0 {
		return fmt.Sprintf("%v", k)
	}

	return fmt.Sprintf("%v.%s", k, strings.Join(fields, "."))
}
