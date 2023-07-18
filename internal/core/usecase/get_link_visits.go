package usecase

import (
	"context"

	"gotiny/internal/core/model"
	"gotiny/internal/core/port"
)

type GetLinkVisits struct {
	repository port.LinksRepository
}

func NewGetLinkVisits(repository port.LinksRepository) *GetLinkVisits {
	return &GetLinkVisits{repository: repository}
}

func (u *GetLinkVisits) Call(
	ctx context.Context,
	linkId string,
	page *string,
) (model.PagedResponse[model.LinkHitAnalitics], error) {
	result, err := u.repository.GetLinkVisits(ctx, linkId, page)
	if err != nil {
		return model.PagedResponse[model.LinkHitAnalitics]{}, &model.AppError{
			Type:    string(model.UnknownError),
			Context: err,
		}
	}
	return result, nil
}
