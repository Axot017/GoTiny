package api

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"

	"gotiny/internal/api/handler"
	"gotiny/internal/api/middleware"
	"gotiny/internal/api/util"
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
		fx.Annotate(
			handler.NewGetLinkDetailsHandler,
			fx.ResultTags(`group:"routes"`),
			fx.As(new(RouteHandler)),
		),
		fx.Annotate(
			handler.NewDeleteLinkHandler,
			fx.ResultTags(`group:"routes"`),
			fx.As(new(RouteHandler)),
		),
		fx.Annotate(
			handler.NewGetVisitsHandler,
			fx.ResultTags(`group:"routes"`),
			fx.As(new(RouteHandler)),
		),
		fx.Annotate(
			handler.NewAjaxHomePageHandler,
			fx.ResultTags(`group:"routes"`),
			fx.As(new(RouteHandler)),
		),
		fx.Annotate(
			handler.NewRedirectHandler,
			fx.ResultTags(`group:"routes"`),
			fx.As(new(RouteHandler)),
		),
		fx.Annotate(
			handler.NewAjaxCreateLinkHandler,
			fx.ResultTags(`group:"routes"`),
			fx.As(new(RouteHandler)),
		),
		fx.Annotate(
			handler.NewAjaxLinkDetailsPageHandler,
			fx.ResultTags(`group:"routes"`),
			fx.As(new(RouteHandler)),
		),
		fx.Annotate(
			handler.NewAjaxDeleteLinkHandler,
			fx.ResultTags(`group:"routes"`),
			fx.As(new(RouteHandler)),
		),
		fx.Annotate(
			handler.NewAjaxGetVisitsHandler,
			fx.ResultTags(`group:"routes"`),
			fx.As(new(RouteHandler)),
		),
		middleware.NewLinkTokenValidator,
		util.NewStructuredLogger,
	}
}
