package usecase

import (
	"context"
)

type DeleteLinkError string

type DeleteLinkRepository interface {
	DeleteLinkById(ctx context.Context, id string) error
}

type DeleteLink struct {
	repository DeleteLinkRepository
}

func NewDeleteLink(repository DeleteLinkRepository) *DeleteLink {
	return &DeleteLink{
		repository: repository,
	}
}

func (u *DeleteLink) Call(ctx context.Context, id string) error {
	return u.repository.DeleteLinkById(ctx, id)
}
