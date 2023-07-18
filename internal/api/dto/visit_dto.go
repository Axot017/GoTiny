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
	// The type of IP address
	//
	// required: false
	Type *string `json:"ip_type"`
	// The country code of the visitor
	//
	// required: false
	CountryCode *string `json:"country_code"`
	// The country name of the visitor
	//
	// required: false
	Country *string `json:"country"`
	// The region name of the visitor
	//
	// required: false
	Region *string `json:"region"`
	// The city name of the visitor
	//
	// required: false
	City *string `json:"city"`
	// The zip code of the visitor
	//
	// required: false
	Zip *string `json:"zip"`
	// The latitude of the visitor
	//
	// required: false
	Latitude *float64 `json:"latitude"`
	// The longitude of the visitor
	//
	// required: false
	Longitude *float64 `json:"longitude"`
	// The time when the visit was created
	//
	// required: true
	CreatedAt time.Time `json:"created_at"`
}

func VisitDtoFromModel(visit model.LinkHitAnalitics) VisitDto {
	if visit.IpDetails == nil {
		return VisitDto{
			Id:        visit.Id,
			IpAddress: visit.RequestData.Ip,
			UserAgent: visit.RequestData.UserAgent,
			CreatedAt: visit.CreatedAt,
		}
	}
	return VisitDto{
		Id:          visit.Id,
		IpAddress:   visit.RequestData.Ip,
		UserAgent:   visit.RequestData.UserAgent,
		Type:        &visit.IpDetails.Type,
		CountryCode: &visit.IpDetails.CountryCode,
		Country:     &visit.IpDetails.Country,
		Region:      &visit.IpDetails.Region,
		City:        &visit.IpDetails.City,
		Zip:         &visit.IpDetails.Zip,
		Latitude:    &visit.IpDetails.Latitude,
		Longitude:   &visit.IpDetails.Longitude,
		CreatedAt:   visit.CreatedAt,
	}
}
