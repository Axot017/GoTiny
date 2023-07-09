package dto

import (
	"time"

	"gotiny/internal/core/model"
)

// swagger:model
type LinkDto struct {
	// Link id
	//
	// required: true
	// example: abc123
	Id string `json:"id"`
	// Short link
	//
	// required: true
	// example: https://{base_url}/abc123
	ShortLink string `json:"short_link"`
	// Original link
	//
	// required: true
	// example: https://google.com
	OriginalLink string `json:"original_link"`
	// Link token - used for link deletion and fetching link details
	//
	// required: true
	Token string `json:"token"`
	// Link visits
	//
	// required: true
	// example: 42
	Hits uint `json:"hits"`
	// Number of link visits allowed, after which the link will be deleted.
	// If not provided, the link will be valid forever
	//
	// required: false
	MaxHits *uint `json:"max_hits,omitempty"`
	// Link expiration date. If not provided, the link will be valid forever
	//
	// required: false
	ValidUntil *time.Time `json:"valid_until,omitempty"`
	// Link creation date
	//
	// required: true
	CreatedAt time.Time `json:"created_at"`

	// Data about link visits will be stored until this date.
	//
	// example: 2021-01-01T00:00:00Z
	// required: false
	TrackUntil *time.Time `json:"track_until" validate:"omitempty"`
}

func LinkDtoFromModel(link model.Link) LinkDto {
	return LinkDto{
		Id:           link.Id,
		ShortLink:    link.ShortLink,
		OriginalLink: link.OriginalLink,
		Token:        link.Token,
		Hits:         link.Hits,
		MaxHits:      link.MaxHits,
		ValidUntil:   link.ValidUntil,
		CreatedAt:    link.CreatedAt,
		TrackUntil:   link.TrackUntil,
	}
}
