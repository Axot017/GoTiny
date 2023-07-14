package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewFromIndex(t *testing.T) {
	index := uint(10)
	url := "https://www.google.com"
	validUntil := time.Now().Add(time.Minute * 60)
	maxHits := uint(10)

	config := LinkToCreate{
		ValidUntil: &validUntil,
		MaxHits:    &maxHits,
		Url:        url,
	}

	link := NewFromIndex(index, config, "http://localhost:8080")

	assert.Equal(t, "a", link.Id)
	assert.Equal(t, "http://localhost:8080/a", link.ShortLink)
	assert.Equal(t, url, link.OriginalLink)
	assert.Equal(t, uint(10), *link.MaxHits)
	assert.Equal(t, uint(0), link.Hits)
	assert.NotNil(t, link.ValidUntil)
	assert.Equal(t, validUntil, *link.ValidUntil)
	assert.NotEmpty(t, link.Token)
}

func TestLinkValidNow(t *testing.T) {
	now := time.Now()
	validUntil := now.Add(time.Minute * 60)

	validLink := Link{
		ValidUntil: &validUntil,
	}

	assert.True(t, validLink.ValidNow())
}

func TestLinkMaxHitsExceeded(t *testing.T) {
	maxHits := uint(10)

	link := Link{
		MaxHits: &maxHits,
		Hits:    10,
	}

	assert.True(t, link.MaxHitsExceeded())
}

func TestLinkValid(t *testing.T) {
	now := time.Now()
	validUntil := now.Add(time.Minute * 60)
	maxHits := uint(10)

	validLink := Link{
		ValidUntil: &validUntil,
		MaxHits:    &maxHits,
		Hits:       9,
	}

	assert.True(t, validLink.Valid())
}
