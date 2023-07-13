package model

import "time"

type LinkConfig struct {
	ValidUntil *time.Time
	MaxHits    *uint
	TrackUntil *time.Time
}
