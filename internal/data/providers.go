package data

import (
	"go.uber.org/fx"

	"gotiny/internal/core/usecase"
	"gotiny/internal/data/adapter"
)

func Providers() []interface{} {
	return []interface{}{
		fx.Annotate(
			adapter.NewLocalLinksRepository,
			fx.As(new(usecase.CreateShortLinkRepository)),
			fx.As(new(usecase.HitLinkRepository)),
		),
	}
}
