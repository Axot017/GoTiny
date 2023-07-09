package model

import "time"

type LinkVisit struct {
	Id        string
	LinkId    string
	IpAddress string
	UserAgent string
	CreatedAt time.Time
}
