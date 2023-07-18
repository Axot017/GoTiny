package port

import (
	"context"

	"gotiny/internal/core/model"
)

type IpRepository interface {
	GetIpDetails(ctx context.Context, ip string) (model.IpDetails, error)
}

type IpCacheRepository interface {
	GetIpDetails(ctx context.Context, ip string) (*model.IpDetails, error)

	SaveIpDetails(ctx context.Context, details model.IpDetails) error
}
