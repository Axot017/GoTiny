package dto

import (
	"time"

	"gotiny/internal/core/model"
)

// swagger:model
type VisitDto struct {
	// The id of the visit
	//
	// required: true
	Id string `json:"id"`
	// Ip address of the visitor
	//
	// required: true
	IpAddress string `json:"ip_addr"`
	// User agent of the visitor
	//
	// required: true
	UserAgent string `json:"user_agent"`
	// The time when the visit was created
	//
	// required: true
	CreatedAt time.Time `json:"created_at"`
}

func VisitDtoFromModel(visit model.LinkVisit) VisitDto {
	return VisitDto{
		Id:        visit.Id,
		IpAddress: visit.IpAddress,
		UserAgent: visit.UserAgent,
		CreatedAt: visit.CreatedAt,
	}
}
