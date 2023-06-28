package adapter

import (
	"context"
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

func (r *LocalLinksRepository) GetNextLinkIndex(ctx context.Context) (uint, error) {
	new := atomic.AddUint64(&r.index, 1)

	return uint(new - 1), nil
}

func (r *LocalLinksRepository) SaveLink(ctx context.Context, link model.Link) error {
	r.links[link.Id] = link

	return nil
}

func (r *LocalLinksRepository) GetLinkById(ctx context.Context, id string) (*model.Link, error) {
	link, ok := r.links[id]
	if ok {
		return &link, nil
	}
	return nil, nil
}

func (r *LocalLinksRepository) DeleteLinkById(ctx context.Context, id string) error {
	delete(r.links, id)
	return nil
}

func (r *LocalLinksRepository) IncrementHitCount(ctx context.Context, id string) error {
	link, ok := r.links[id]
	if ok {
		link.Hits++
		r.links[id] = link
	}
	return nil
}
