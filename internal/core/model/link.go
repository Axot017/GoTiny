package model

import (
	"time"

	"gotiny/internal/core/util"
)

type Link struct {
	Id           string
	ShortLink    string
	OriginalLink string
	Hits         uint
	MaxHits      *uint
	ValidUntil   *time.Time
	CreatedAt    time.Time
}

func NewFromIndex(index uint, url string, config LinkConfig, baseUrl string) Link {
	id := util.EncodeBase62(index)
	shortLink := baseUrl + "/" + id
	now := time.Now()

	return Link{
		Id:           id,
		ShortLink:    shortLink,
		OriginalLink: url,
		CreatedAt:    now,
		MaxHits:      config.MaxHits,
		ValidUntil:   config.ValidUntil,
	}
}

func (l Link) Valid() bool {
	return !l.MaxHitsExceeded() && l.ValidNow()
}

func (l Link) MaxHitsExceeded() bool {
	return l.MaxHits != nil && l.Hits >= *l.MaxHits
}

func (l Link) ValidNow() bool {
	return l.ValidUntil != nil && l.ValidUntil.After(time.Now())
}
