package usecase

import "gotiny/internal/core/model"

type HitLinkRepository interface {
	GetLinkById(id string) (*model.Link, error)

	DeleteLinkById(id string) error

	IncrementHitCount(id string) error
}

type HitLink struct {
	repository HitLinkRepository
}

func NewHitLink(repository HitLinkRepository) *HitLink {
	return &HitLink{
		repository: repository,
	}
}

func (u *HitLink) Call(id string) (*string, error) {
	link, err := u.repository.GetLinkById(id)
	if err != nil {
		return nil, err
	}

	if link == nil {
		return nil, nil
	}

	if link.Valid() {
		go u.repository.DeleteLinkById(id)
		return nil, nil
	}

	go u.repository.IncrementHitCount(id)

	return &link.OriginalLink, nil
}
