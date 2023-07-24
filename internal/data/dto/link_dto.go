package dto

import (
	"time"

	"gotiny/internal/core/model"
)

const (
	LinkPK           = "LINK"
	LinkSKPrefix     = "LINK#"
	LinkGSI1PKPrefix = "USER#"
	LinkGSI1SKPrefix = "LINK#"
)

type LinkDto struct {
	PK                      string // Constant
	SK                      string // Link id
	ShortLink               string
	OriginalLink            string
	Token                   string
	Hits                    uint
	MaxHits                 *uint   `dynamodbav:",omitempty"`
	TTL                     *uint   `dynamodbav:",omitempty"`
	GSI_1_PK                *string `dynamodbav:",omitempty"` // User id
	GSI_1_SK                *string `dynamodbav:",omitempty"` // Created at
	EnableDetailedAnalytics bool
	CreatedAt               time.Time
}

func LinkDtoFromLink(link model.Link) LinkDto {
	var ttl *uint
	if link.ValidUntil != nil {
		unix := uint(link.ValidUntil.Unix())
		ttl = &unix
	}
	var gsi1PK *string
	var gsi1SK *string
	if link.UserId != nil {
		pk := LinkGSI1PKPrefix + *link.UserId
		sk := LinkGSI1SKPrefix + link.CreatedAt.Format(time.RFC3339Nano)
		gsi1PK = &pk
		gsi1SK = &sk
	}
	return LinkDto{
		PK:                      LinkPK,
		SK:                      LinkSKPrefix + link.Id,
		ShortLink:               link.ShortLink,
		OriginalLink:            link.OriginalLink,
		Token:                   link.Token,
		Hits:                    link.Hits,
		MaxHits:                 link.MaxHits,
		TTL:                     ttl,
		GSI_1_PK:                gsi1PK,
		GSI_1_SK:                gsi1SK,
		EnableDetailedAnalytics: link.EnableDetailedAnalytics,
		CreatedAt:               link.CreatedAt,
	}
}

func LinkDtoToLink(d LinkDto) model.Link {
	var validUntil *time.Time
	if d.TTL != nil {
		t := time.Unix(int64(*d.TTL), 0)
		validUntil = &t
	}
	var userId *string
	if d.GSI_1_PK != nil {
		u := (*d.GSI_1_PK)[len(LinkGSI1PKPrefix):]
		userId = &u
	}
	return model.Link{
		Id:                      d.SK[len(LinkSKPrefix):],
		ShortLink:               d.ShortLink,
		OriginalLink:            d.OriginalLink,
		Token:                   d.Token,
		Hits:                    d.Hits,
		MaxHits:                 d.MaxHits,
		ValidUntil:              validUntil,
		UserId:                  userId,
		CreatedAt:               d.CreatedAt,
		EnableDetailedAnalytics: d.EnableDetailedAnalytics,
	}
}
