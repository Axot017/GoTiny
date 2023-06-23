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

func (u *CreateShortLink) Call(url string, link_config model.LinkConfig) (model.Link, error) {
	index, err := u.repository.GetNextLinkIndex()
	if err != nil {
		return model.Link{}, err
	}

	link := model.NewFromIndex(index, url, link_config)

	err = u.repository.SaveLink(link)
	if err != nil {
		return model.Link{}, err
	}

	return link, nil
}
