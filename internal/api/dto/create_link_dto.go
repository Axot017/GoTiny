package dto

type CreateLinkDto struct {
	Link     string `json:"link"       validate:"required,http_url"`
	TtlInSec *uint  `json:"ttl_in_sec" validate:"omitempty,gte=0"`
	MaxHits  *uint  `json:"max_hits"   validate:"omitempty,gte=0"`
}
