package usecase

import (
	"errors"

	"gotiny/internal/core/model"
)

type DeleteLinkRepository interface {
	GetLinkById(id string) (*model.Link, error)

	DeleteLinkById(id string) error
}

type DeleteLink struct {
	repository DeleteLinkRepository
}

func NewDeleteLink(repository DeleteLinkRepository) *DeleteLink {
	return &DeleteLink{
		repository: repository,
	}
}

func (u *DeleteLink) Call(id string, token string) error {
	link, err := u.repository.GetLinkById(id)
	if err != nil {
		return err
	}

	if link == nil {
		return errors.New("link not found")
	}

	if link.Token != token {
		return errors.New("invalid token")
	}

	return u.repository.DeleteLinkById(id)
}
