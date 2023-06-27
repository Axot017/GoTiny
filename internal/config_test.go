package internal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseUrl(t *testing.T) {
	os.Setenv("BASE_URL", "https://www.google.com")
	os.Setenv("LOG_JSON", "true")
	config := NewConfig()
	assert.Equal(t, "https://www.google.com", config.BaseUrl())
	assert.Equal(t, true, config.LogJson())
}
