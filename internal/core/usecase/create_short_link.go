package usecase

import (
	"gotiny/internal/core/model"
	"gotiny/internal/core/port"
)

type CreateShortLink struct {
	repository port.LinksRepository
}

func NewCreateShortLink(repository port.LinksRepository) *CreateShortLink {
	return &CreateShortLink{repository}
}

func (u *CreateShortLink) Call(url string) (model.Link, error) {
	index, err := u.repository.GetNextLinkIndex()
	if err != nil {
		// TODO: implove
		return model.Link{}, err
	}
	return model.Link{}, nil
}
