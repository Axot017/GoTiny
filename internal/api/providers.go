package api

import (
	"go.uber.org/fx"

	"gotiny/internal/api/handler"
)

func Providers() []interface{} {
	return []interface{}{
		fx.Annotate(
			handler.NewHealthHandler,
			fx.ResultTags(`group:"routes"`),
			fx.As(new(RouteHandler)),
		),
		fx.Annotate(
			handler.NewCreateLinkHandler,
			fx.ResultTags(`group:"routes"`),
			fx.As(new(RouteHandler)),
		),
	}
}
