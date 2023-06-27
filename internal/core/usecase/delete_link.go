package usecase

import (
	"gotiny/internal/core/model"
)

type DeleteLinkError string

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
		return &model.AppError{
			Type:    string(model.UnknownError),
			Context: err,
		}
	}

	if link == nil {
		return &model.AppError{
			Type: string(model.NotFoundError),
		}
	}

	if link.Token != token {
		return &model.AppError{
			Type: string(model.UnauthorizedError),
		}
	}

	return u.repository.DeleteLinkById(id)
}
