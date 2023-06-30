package usecase

import (
	"context"

	"gotiny/internal/core/model"
)

type HitLinkRepository interface {
	GetLinkById(ctx context.Context, id string) (*model.Link, error)

	DeleteLinkById(ctx context.Context, id string) error

	IncrementHitCount(ctx context.Context, id string) error
}

type HitLink struct {
	repository HitLinkRepository
}

func NewHitLink(repository HitLinkRepository) *HitLink {
	return &HitLink{
		repository: repository,
	}
}

func (u *HitLink) Call(ctx context.Context, id string) (*string, error) {
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

	if !link.Valid() {
		go u.repository.DeleteLinkById(context.Background(), id)
		return nil, nil
	}

	go u.repository.IncrementHitCount(context.Background(), id)

	return &link.OriginalLink, nil
}
