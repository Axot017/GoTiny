package model

import (
	"time"

	"gotiny/internal/core/util"
)

type Link struct {
	Id           string
	ShortLink    string
	OriginalLink string
	Hits         int
	MaxHits      *int
	ValidUntil   *time.Time
	CreatedAt    time.Time
}

func NewFromIndex(index uint, url string, config LinkConfig) Link {
	id := util.EncodeBase62(index)
	shortLink := config.Protocol + "://" + config.Host + "/" + id
	now := time.Now()

	var validUntil *time.Time
	if config.TtlInSec != nil {
		time := now.Add(time.Duration(*config.TtlInSec) * time.Second)
		validUntil = &time
	}

	return Link{
		Id:           id,
		ShortLink:    shortLink,
		OriginalLink: url,
		CreatedAt:    now,
		MaxHits:      config.MaxHits,
		ValidUntil:   validUntil,
	}
}
