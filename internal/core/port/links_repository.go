package port

import "gotiny/internal/core/model"

type LinksRepository interface {
	GetNextLinkIndex() (uint, error)

	SaveLink(link model.Link) error
}
