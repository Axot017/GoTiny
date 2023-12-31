package core

import (
	"os"

	"golang.org/x/exp/slog"

	"gotiny/internal/core/port"
)

func NewSlogHandler(config port.Config) slog.Handler {
	opitons := slog.HandlerOptions{}
	var handler slog.Handler

	if config.LogJson() {
		handler = slog.NewJSONHandler(os.Stdout, &opitons)
	} else {
		handler = slog.NewTextHandler(os.Stdout, &opitons)
	}

	slog.SetDefault(slog.New(handler))

	return handler
}
