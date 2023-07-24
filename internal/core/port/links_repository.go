package port

import (
	"context"

	"gotiny/internal/core/model"
)

type LinksRepository interface {
	GetLinkById(ctx context.Context, id string) (*model.Link, error)

	DeleteLinkById(ctx context.Context, id string) error

	IncrementHitCount(ctx context.Context, id string) error

	SaveHitAnalitics(
		ctx context.Context,
		linkId string,
		requestData model.LinkHitAnalitics,
	) error

	GetNextLinkIndex(ctx context.Context) (uint, error)

	SaveLink(ctx context.Context, link model.Link) error

	GetLinkVisits(
		ctx context.Context,
		linkId string,
		page *string,
	) (model.PagedResponse[model.LinkHitAnalitics], error)

	GetUserLinks(
		ctx context.Context,
		userId string,
		page *string,
	) (model.PagedResponse[model.Link], error)
}
