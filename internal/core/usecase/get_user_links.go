package usecase

import (
	"context"

	"gotiny/internal/core/model"
	"gotiny/internal/core/port"
)

type GetUserLinks struct {
	repository port.LinksRepository
}

func NewGetUserLinks(repository port.LinksRepository) *GetUserLinks {
	return &GetUserLinks{repository}
}

func (u *GetUserLinks) Call(
	ctx context.Context,
	userId string,
	page *string,
) (model.PagedResponse[model.Link], error) {
	links, err := u.repository.GetUserLinks(ctx, userId, page)
	if err != nil {
		return model.PagedResponse[model.Link]{}, &model.AppError{
			Type:    string(model.UnknownError),
			Context: err,
		}
	}

	return links, nil
}
