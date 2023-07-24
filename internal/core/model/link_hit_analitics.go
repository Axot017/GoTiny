package model

import "time"

type LinkHitAnalitics struct {
	Id          string
	RequestData RedirectRequestData
	IpDetails   *IpDetails
	CreatedAt   time.Time
}
