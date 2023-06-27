package core

import (
	"go.uber.org/fx"
	"golang.org/x/exp/slog"

	"gotiny/internal/core/usecase"
)

func Providers() []interface{} {
	return []interface{}{
		usecase.NewCreateShortLink,
		usecase.NewHitLink,
		usecase.NewGetLinkDetails,
		usecase.NewDeleteLink,
		fx.Annotate(NewSlogHandler, fx.As(new(slog.Handler))),
	}
}
