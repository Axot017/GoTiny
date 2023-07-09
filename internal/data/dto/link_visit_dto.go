package dto

import (
	"time"

	"gotiny/internal/core/model"
)

const (
	LinkVisitPKPrefix = "LINK#"
	LinkVisitSKPrefix = "VISIT#"
)

type LinkVisitDto struct {
	PK        string
	SK        string
	Ip        string
	UserAgent string
	CreatedAt time.Time
	TTL       *uint
	// TODO: location
}

func LinkVisitDtoToModel(dto LinkVisitDto) model.LinkVisit {
	return model.LinkVisit{
		Id:        dto.SK[len(LinkVisitSKPrefix):],
		IpAddress: dto.Ip,
		UserAgent: dto.UserAgent,
		CreatedAt: dto.CreatedAt,
	}
}
