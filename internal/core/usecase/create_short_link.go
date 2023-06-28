package usecase

import (
	"context"

	"gotiny/internal/core/model"
)

type CreateShortLinkConfig interface {
	BaseUrl() string
}

type CreateShortLinkRepository interface {
	GetNextLinkIndex(ctx context.Context) (uint, error)

	SaveLink(ctx context.Context, link model.Link) error
}

type CreateShortLink struct {
	repository CreateShortLinkRepository
	config     CreateShortLinkConfig
}

func NewCreateShortLink(
	repository CreateShortLinkRepository,
	config CreateShortLinkConfig,
) *CreateShortLink {
	return &CreateShortLink{repository, config}
}

func (u *CreateShortLink) Call(
	ctx context.Context,
	url string,
	link_config model.LinkConfig,
) (model.Link, error) {
	index, err := u.repository.GetNextLinkIndex(ctx)
	if err != nil {
		return model.Link{}, &model.AppError{
			Type:    string(model.UnknownError),
			Context: err,
		}
	}

	link := model.NewFromIndex(index, url, link_config, u.config.BaseUrl())

	err = u.repository.SaveLink(ctx, link)
	if err != nil {
		return model.Link{}, &model.AppError{
			Type:    string(model.UnknownError),
			Context: err,
		}
	}

	return link, nil
}
