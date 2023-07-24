package model

import "time"

type LinkToCreate struct {
	Url                     string
	ValidUntil              *time.Time
	MaxHits                 *uint
	EnableDetailedAnalytics bool
	UserId                  *string
}
