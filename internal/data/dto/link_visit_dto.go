package dto

import "time"

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
