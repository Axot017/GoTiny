package adapter

import (
	"sync/atomic"

	"gotiny/internal/core/model"
)

type LocalLinksRepository struct {
	links map[string]model.Link
	index uint64
}

func NewLocalLinksRepository() *LocalLinksRepository {
	return &LocalLinksRepository{
		links: make(map[string]model.Link),
	}
}

func (r *LocalLinksRepository) GetNextLinkIndex() (uint, error) {
	new := atomic.AddUint64(&r.index, 1)

	return uint(new - 1), nil
}

func (r *LocalLinksRepository) SaveLink(link model.Link) error {
	r.links[link.Id] = link

	return nil
}
