package dto

import "time"

// swagger:model
type CreateLinkDto struct {
	// Link to be shortened
	//
	// required: true
	// example: https://google.com
	// format: url
	Link string `json:"link"        validate:"required,url"`
	// Valid until in iso8601 format. If not provided, the link will be valid forever
	//
	// example: 2021-01-01T00:00:00Z
	// required: false
	ValidUntil *time.Time `json:"valid_until" validate:"omitempty"`
	// Max link visits. If not provided, the link will be valid forever
	// example: 10
	// required: false
	// minimum: 1
	MaxHits *uint `json:"max_hits"    validate:"omitempty,gt=0"`

	// Data about link visits will be stored until this date.
	// If not provided, the data will be not stored
	//
	// example: 2021-01-01T00:00:00Z
	// required: false
	TrackUntil *time.Time `json:"track_until" validate:"omitempty"`
}
