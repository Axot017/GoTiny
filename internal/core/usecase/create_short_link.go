package usecase

import (
	"context"
	"net/url"
	"strings"
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
	if !strings.HasPrefix(linkToCreate.Url, "http") {
		linkToCreate.Url = "https://" + linkToCreate.Url
	}
	isValid := u.isValidUrl(ctx, linkToCreate.Url)

	if !isValid {
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

func (u *CreateShortLink) isValidUrl(ctx context.Context, link string) bool {
	_, err := url.ParseRequestURI(link)
	if err != nil {
		return false
	}

	isValid, err := u.urlChecker.CheckUrl(ctx, link)
	if err != nil {
		return false
	}
	return isValid
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
