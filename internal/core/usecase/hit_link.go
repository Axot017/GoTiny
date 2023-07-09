package usecase

import (
	"context"
	"time"

	"gotiny/internal/core/model"
)

type HitLinkRepository interface {
	GetLinkById(ctx context.Context, id string) (*model.Link, error)

	DeleteLinkById(ctx context.Context, id string) error

	IncrementHitCount(ctx context.Context, id string) error

	SaveRedirectRequestData(
		ctx context.Context,
		linkId string,
		requestData model.RedirecsRequestData,
	) error
}

type HitLink struct {
	repository HitLinkRepository
}

func NewHitLink(repository HitLinkRepository) *HitLink {
	return &HitLink{
		repository: repository,
	}
}

func (u *HitLink) Call(
	ctx context.Context,
	id string,
	requestData model.RedirecsRequestData,
) (*string, error) {
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

	go u.saveRequestData(link, requestData)

	if !link.Valid() {
		go u.repository.DeleteLinkById(context.Background(), id)
		return nil, nil
	}

	go u.repository.IncrementHitCount(context.Background(), id)

	return &link.OriginalLink, nil
}

func (u *HitLink) saveRequestData(link *model.Link, requestData model.RedirecsRequestData) {
	if link.TrackUntil == nil {
		return
	}

	if link.TrackUntil.Before(time.Now()) {
		return
	}

	u.repository.SaveRedirectRequestData(context.Background(), link.Id, requestData)
}
