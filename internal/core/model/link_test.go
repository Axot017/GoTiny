package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFromIndex(t *testing.T) {
	index := uint(10)
	url := "https://www.google.com"
	ttl := uint(60)
	maxHits := uint(10)

	config := LinkConfig{
		TtlInSec: &ttl,
		Host:     "localhost:8080",
		MaxHits:  &maxHits,
	}

	link := NewFromIndex(index, url, config, "http://localhost:8080")

	assert.Equal(t, "a", link.Id)
	assert.Equal(t, "http://localhost:8080/a", link.ShortLink)
	assert.Equal(t, url, link.OriginalLink)
	assert.Equal(t, uint(10), *link.MaxHits)
	assert.Equal(t, uint(0), link.Hits)
	assert.NotNil(t, link.ValidUntil)
	assert.Equal(t, 60, int(link.ValidUntil.Sub(link.CreatedAt).Seconds()))
}
