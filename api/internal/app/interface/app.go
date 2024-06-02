package app

import (
	isql "webapi/pkg/postgres/interface"

	"golang.org/x/exp/slog"
)

type App interface {
	Service() Service
	Storage() Storage
	Logger() *slog.Logger
	DB() isql.Runner
	Validator() Validator
}
