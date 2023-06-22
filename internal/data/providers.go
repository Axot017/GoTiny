package data

import (
	"go.uber.org/fx"

	"gotiny/internal/core/port"
	"gotiny/internal/data/adapter"
)

func Providers() []interface{} {
	return []interface{}{
		fx.Annotate(adapter.NewLocalLinksRepository, fx.As(new(port.LinksRepository))),
	}
}
