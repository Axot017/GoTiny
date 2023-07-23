package model

import "time"

type LinkToCreate struct {
	Url        string
	ValidUntil *time.Time
	MaxHits    *uint
	TrackUntil *time.Time
	UserId     *string
}
