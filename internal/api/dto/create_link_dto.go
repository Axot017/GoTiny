package dto

import "time"

type CreateLinkDto struct {
	Link       string     `json:"link"        validate:"required,url"`
	ValidUntil *time.Time `json:"valid_until" validate:"omitempty"`
	MaxHits    *uint      `json:"max_hits"    validate:"omitempty,gt=0"`
}
