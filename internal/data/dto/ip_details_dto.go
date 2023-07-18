package dto

import "gotiny/internal/core/model"

const (
	IpPK       = "IP"
	IpSKPrefix = "IP#"
)

type IpDetailsDto struct {
	PK          string
	SK          string
	Type        string
	CountryCode string
	Country     string
	Region      string
	City        string
	Zip         string
	Latitude    float64
	Longitude   float64
}

func (r *IpDetailsDto) ToIpDetails() model.IpDetails {
	return model.IpDetails{
		Ip:          r.SK[len(IpSKPrefix):],
		Type:        r.Type,
		CountryCode: r.CountryCode,
		Country:     r.Country,
		Region:      r.Region,
		City:        r.City,
		Zip:         r.Zip,
		Latitude:    r.Latitude,
		Longitude:   r.Longitude,
	}
}

func IpDetailsDtoFromModel(m model.IpDetails) IpDetailsDto {
	return IpDetailsDto{
		PK:          IpPK,
		SK:          IpSKPrefix + m.Ip,
		Type:        m.Type,
		CountryCode: m.CountryCode,
		Country:     m.Country,
		Region:      m.Region,
		City:        m.City,
		Zip:         m.Zip,
		Latitude:    m.Latitude,
		Longitude:   m.Longitude,
	}
}
