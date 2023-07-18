package model

import "time"

type LinkHitAnalitics struct {
	Id          string
	RequestData RedirecsRequestData
	IpDetails   *IpDetails
	CreatedAt   time.Time
}
