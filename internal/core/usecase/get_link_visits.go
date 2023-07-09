package usecase

import (
	"context"

	"gotiny/internal/core/model"
)

type GetLinkVisitsRepository interface {
	GetLinkVisits(
		ctx context.Context,
		linkId string,
		page *string,
	) (model.PagedResponse[model.LinkVisit], error)
}

type GetLinkVisits struct {
	repository GetLinkVisitsRepository
}

func NewGetLinkVisits(repository GetLinkVisitsRepository) *GetLinkVisits {
	return &GetLinkVisits{repository: repository}
}

func (u *GetLinkVisits) Call(
	ctx context.Context,
	linkId string,
	page *string,
) (model.PagedResponse[model.LinkVisit], error) {
	result, err := u.repository.GetLinkVisits(ctx, linkId, page)
	if err != nil {
		return model.PagedResponse[model.LinkVisit]{}, &model.AppError{
			Type:    string(model.UnknownError),
			Context: err,
		}
	}
	return result, nil
}
