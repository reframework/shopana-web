package application

import (
	"os"

	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
)

func CreateLogger() *slog.Logger {
	logType := viper.GetString("LOG_TYPE")

	var handler slog.Handler

	switch logType {
	case "json":
		handler = slog.NewJSONHandler(os.Stderr, nil)
	default: // Default to text logger
		handler = slog.NewTextHandler(os.Stderr, nil)
	}

	return slog.New(handler)
}
