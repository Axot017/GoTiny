package usecase

import (
	"context"
	"time"

	"gotiny/internal/core/model"
	"gotiny/internal/core/port"
)

const (
	InvalidUrlError = "invalid_url"
)

type CreateShortLink struct {
	repository port.LinksRepository
	config     port.Config
	urlChecker port.LinkChecker
}

func NewCreateShortLink(
	repository port.LinksRepository,
	config port.Config,
	urlChecker port.LinkChecker,
) *CreateShortLink {
	return &CreateShortLink{repository, config, urlChecker}
}

func (u *CreateShortLink) Call(
	ctx context.Context,
	linkToCreate model.LinkToCreate,
) (model.Link, error) {
	isValid, err := u.urlChecker.CheckUrl(ctx, linkToCreate.Url)

	if err == nil && !isValid {
		return model.Link{}, &model.AppError{
			Type: InvalidUrlError,
		}
	}

	index, err := u.repository.GetNextLinkIndex(ctx)
	if err != nil {
		return model.Link{}, &model.AppError{
			Type:    string(model.UnknownError),
			Context: err,
		}
	}

	link := model.NewFromIndex(index, linkToCreate, u.config.BaseUrl())
	link.TrackUntil = u.getTrackUntilDate(link)

	err = u.repository.SaveLink(ctx, link)
	if err != nil {
		return model.Link{}, &model.AppError{
			Type:    string(model.UnknownError),
			Context: err,
		}
	}

	return link, nil
}

func (u *CreateShortLink) getTrackUntilDate(link model.Link) *time.Time {
	if link.TrackUntil == nil {
		return nil
	}
	maxTrackingDays := u.config.MaxTrackingDays()
	now := time.Now()
	maxTrackingDate := now.AddDate(0, 0, int(maxTrackingDays))
	if link.TrackUntil.After(maxTrackingDate) {
		return &maxTrackingDate
	} else {
		return link.TrackUntil
	}
}
