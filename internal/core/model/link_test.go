package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFromIndex(t *testing.T) {
	index := uint(10)
	url := "https://www.google.com"
	ttl := 60
	maxHits := 10

	config := LinkConfig{
		TtlInSec: &ttl,
		Host:     "localhost:8080",
		Protocol: "http",
		MaxHits:  &maxHits,
	}

	link := NewFromIndex(index, url, config)

	assert.Equal(t, "a", link.Id)
	assert.Equal(t, "http://localhost:8080/a", link.ShortLink)
	assert.Equal(t, url, link.OriginalLink)
	assert.Equal(t, 10, *link.MaxHits)
	assert.Equal(t, 0, link.Hits)
	assert.NotNil(t, link.ValidUntil)
	assert.Equal(t, 60, int(link.ValidUntil.Sub(link.CreatedAt).Seconds()))
}
