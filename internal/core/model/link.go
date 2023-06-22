package model

import "time"

type Link struct {
	Id           string    `json:"id"`
	ShortLink    string    `json:"short_link"`
	OriginalLink string    `json:"original_link"`
	Hits         int       `json:"hits"`
	CreatedAt    time.Time `json:"created_at"`
}
