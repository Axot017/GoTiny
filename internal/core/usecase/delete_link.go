package usecase

import (
	"context"

	"gotiny/internal/core/port"
)

type DeleteLink struct {
	repository port.LinksRepository
}

func NewDeleteLink(repository port.LinksRepository) *DeleteLink {
	return &DeleteLink{
		repository: repository,
	}
}

func (u *DeleteLink) Call(ctx context.Context, id string) error {
	return u.repository.DeleteLinkById(ctx, id)
}
