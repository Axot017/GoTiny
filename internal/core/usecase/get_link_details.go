package usecase

import (
	"context"

	"gotiny/internal/core/model"
	"gotiny/internal/core/port"
)

type GetLinkDetails struct {
	repository port.LinksRepository
}

func NewGetLinkDetails(repository port.LinksRepository) *GetLinkDetails {
	return &GetLinkDetails{
		repository: repository,
	}
}

func (u *GetLinkDetails) Call(ctx context.Context, id string, token string) (*model.Link, error) {
	link, err := u.repository.GetLinkById(ctx, id)
	if err != nil {
		return nil, &model.AppError{
			Type:    string(model.UnknownError),
			Context: err,
		}
	}
	if link == nil {
		return nil, nil
	}
	if link.Token != token {
		return nil, &model.AppError{
			Type: string(model.UnauthorizedError),
		}
	}

	return link, nil
}
