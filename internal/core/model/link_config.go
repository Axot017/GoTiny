package model

type LinkConfig struct {
	TtlInSec *int
	MaxHits  *int
	Host     string
	Protocol string
}
