package dto

import (
	"time"

	"gotiny/internal/core/model"
)

const (
	LinkVisitPKPrefix = "LINK#"
	LinkVisitSKPrefix = "VISIT#"
)

type LinkHitAnaliticsDto struct {
	PK          string
	SK          string
	RequestData model.RedirectRequestData
	IpDetails   *model.IpDetails `dynamodbav:",omitempty"`
	CreatedAt   time.Time
	TTL         *uint `dynamodbav:",omitempty"`
}

func LinkHitAnaliticsDtoToDomain(links LinkHitAnaliticsDto) model.LinkHitAnalitics {
	return model.LinkHitAnalitics{
		Id:          links.SK[len(LinkVisitSKPrefix):],
		RequestData: links.RequestData,
		IpDetails:   links.IpDetails,
		CreatedAt:   links.CreatedAt,
	}
}
