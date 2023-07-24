package dto

import "time"

// swagger:model
type CreateLinkDto struct {
	// Link to be shortened
	//
	// required: true
	// example: https://google.com
	// format: url
	Link string `json:"link"        form:"link"        validate:"required"`
	// Valid until in iso8601 format. If not provided, the link will be valid forever
	//
	// example: 2021-01-01T00:00:00Z
	// required: false
	ValidUntil *time.Time `json:"valid_until" form:"valid_until" validate:"omitempty"`
	// Max link visits. If not provided, the link will be valid forever
	// example: 10
	// required: false
	// minimum: 1
	MaxHits *uint `json:"max_hits"    form:"max_hits"    validate:"omitempty,gt=0"`

	// Enable detailed analytics
	// example: true
	EnableDetailedAnalytics bool `json:"enable_detailed_analytics" form:"enable_detailed_analytics" validate:"omitempty"`
}
