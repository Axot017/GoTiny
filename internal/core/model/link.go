package model

import (
	"time"

	"gotiny/internal/core/util"
)

const tokenLength = 16

type Link struct {
	Id                      string
	ShortLink               string
	OriginalLink            string
	Token                   string
	Hits                    uint
	MaxHits                 *uint
	UserId                  *string
	ValidUntil              *time.Time
	EnableDetailedAnalytics bool
	CreatedAt               time.Time
}

func NewFromIndex(index uint, linkToCreate LinkToCreate, baseUrl string) Link {
	id := util.EncodeBase62(index)
	shortLink := baseUrl + "/" + id
	now := time.Now()

	return Link{
		Id:                      id,
		ShortLink:               shortLink,
		Token:                   util.RandString(tokenLength),
		OriginalLink:            linkToCreate.Url,
		CreatedAt:               now,
		UserId:                  linkToCreate.UserId,
		MaxHits:                 linkToCreate.MaxHits,
		EnableDetailedAnalytics: linkToCreate.EnableDetailedAnalytics,
		ValidUntil:              linkToCreate.ValidUntil,
	}
}

func (l Link) Valid() bool {
	return !l.MaxHitsExceeded() && l.ValidNow()
}

func (l Link) MaxHitsExceeded() bool {
	return l.MaxHits != nil && l.Hits >= *l.MaxHits
}

func (l Link) ValidNow() bool {
	return l.ValidUntil == nil || l.ValidUntil.After(time.Now())
}
