package dto

import (
	"time"

	"gotiny/internal/core/model"
)

const (
	LinkPK       = "LINK"
	LinkSKPrefix = "LINK#"
)

type LinkDto struct {
	PK           string
	SK           string
	ShortLink    string
	OriginalLink string
	Token        string
	Hits         uint
	MaxHits      *uint
	TTL          *uint
	TrackUntil   *time.Time
	CreatedAt    time.Time
}

func LinkDtoFromLink(link model.Link) LinkDto {
	var ttl *uint
	if link.ValidUntil != nil {
		unix := uint(link.ValidUntil.Unix())
		ttl = &unix
	}
	return LinkDto{
		PK:           LinkPK,
		SK:           LinkSKPrefix + link.Id,
		ShortLink:    link.ShortLink,
		OriginalLink: link.OriginalLink,
		Token:        link.Token,
		Hits:         link.Hits,
		MaxHits:      link.MaxHits,
		TTL:          ttl,
		TrackUntil:   link.TrackUntil,
		CreatedAt:    link.CreatedAt,
	}
}

func (d LinkDto) ToLink() model.Link {
	var validUntil *time.Time
	if d.TTL != nil {
		t := time.Unix(int64(*d.TTL), 0)
		validUntil = &t
	}
	return model.Link{
		Id:           d.SK[len(LinkSKPrefix):],
		ShortLink:    d.ShortLink,
		OriginalLink: d.OriginalLink,
		Token:        d.Token,
		Hits:         d.Hits,
		MaxHits:      d.MaxHits,
		ValidUntil:   validUntil,
		CreatedAt:    d.CreatedAt,
		TrackUntil:   d.TrackUntil,
	}
}
