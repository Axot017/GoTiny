package dto

import "gotiny/internal/core/model"

type IpStackResponseDto struct {
	Ip            string        `json:"ip"`
	Hostname      string        `json:"hostname"`
	Type          string        `json:"type"`
	ContinentCode string        `json:"continent_code"`
	ContinentName string        `json:"continent_name"`
	CountryCode   string        `json:"country_code"`
	CountryName   string        `json:"country_name"`
	RegionCode    string        `json:"region_code"`
	RegionName    string        `json:"region_name"`
	City          string        `json:"city"`
	Zip           string        `json:"zip"`
	Latitude      float64       `json:"latitude"`
	Longitude     float64       `json:"longitude"`
	Location      LocationDto   `json:"location"`
	TimeZone      TimeZoneDto   `json:"time_zone"`
	Currency      CurrencyDto   `json:"currency"`
	Connection    ConnectionDto `json:"connection"`
	Security      SecurityDto   `json:"security"`
}

type LocationDto struct {
	GeonameId               int           `json:"geoname_id"`
	Capital                 string        `json:"capital"`
	Languages               []LanguageDto `json:"languages"`
	CountryFlag             string        `json:"country_flag"`
	CountryFlagEmoji        string        `json:"country_flag_emoji"`
	CountryFlagEmojiUnicode string        `json:"country_flag_emoji_unicode"`
	CallingCode             string        `json:"calling_code"`
	IsEu                    bool          `json:"is_eu"`
}

type LanguageDto struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Native string `json:"native"`
}

type TimeZoneDto struct {
	Id               string `json:"id"`
	CurrentTime      string `json:"current_time"`
	GmtOffset        int    `json:"gmt_offset"`
	Code             string `json:"code"`
	IsDaylightSaving bool   `json:"is_daylight_saving"`
}

type CurrencyDto struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Plural       string `json:"plural"`
	Symbol       string `json:"symbol"`
	SymbolNative string `json:"symbol_native"`
}

type ConnectionDto struct {
	Asn int    `json:"asn"`
	Isp string `json:"isp"`
}

type SecurityDto struct {
	IsProxy     bool    `json:"is_proxy"`
	ProxyType   *string `json:"proxy_type"`
	IsCrawler   bool    `json:"is_crawler"`
	CrawlerName *string `json:"crawler_name"`
	CrawlerType *string `json:"crawler_type"`
	IsTor       bool    `json:"is_tor"`
	ThreatLevel string  `json:"threat_level"`
}

func (r *IpStackResponseDto) ToIpDetails() model.IpDetails {
	return model.IpDetails{
		Ip:                  r.Ip,
		Type:                r.Type,
		CountryFlagSvgImage: r.Location.CountryFlag,
		CountryCode:         r.CountryCode,
		Country:             r.CountryName,
		Region:              r.RegionName,
		City:                r.City,
		Zip:                 r.Zip,
		Latitude:            r.Latitude,
		Longitude:           r.Longitude,
	}
}
