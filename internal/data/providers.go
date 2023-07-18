package data

import (
	"go.uber.org/fx"

	"gotiny/internal/core/port"
	"gotiny/internal/data/adapter"
)

func Providers() []interface{} {
	return []interface{}{
		fx.Annotate(
			adapter.NewDynamodbLinksRepository,
			fx.As(new(port.LinksRepository)),
		),
		fx.Annotate(
			adapter.NewDynamodIpRepository,
			fx.As(new(port.IpCacheRepository)),
		),
		fx.Annotate(
			adapter.NewIpStackApiClient,
			fx.As(new(port.IpRepository)),
		),
		fx.Annotate(
			adapter.NewHttpClient,
			fx.As(new(port.LinkChecker)),
		),
		newAwsConfig,
		newDynamobdClient,
	}
}
