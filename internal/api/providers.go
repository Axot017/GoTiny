package api

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"

	"gotiny/internal/api/handler"
)

func Providers() []interface{} {
	return []interface{}{
		validator.New,
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
