package data

import (
	"go.uber.org/fx"

	"gotiny/internal/core/usecase"
	"gotiny/internal/data/adapter"
)

func Providers() []interface{} {
	return []interface{}{
		fx.Annotate(
			adapter.NewDynamodbLinksRepository,
			fx.As(new(usecase.CreateShortLinkRepository)),
			fx.As(new(usecase.HitLinkRepository)),
			fx.As(new(usecase.GetLinkDetailsRepository)),
			fx.As(new(usecase.DeleteLinkRepository)),
		),
		newAwsConfig,
		newDynamobdClient,
	}
}
