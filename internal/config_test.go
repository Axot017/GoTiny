package internal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseUrl(t *testing.T) {
	os.Setenv("BASE_URL", "https://www.google.com")
	config := NewConfig()
	assert.Equal(t, "https://www.google.com", config.BaseUrl())
}
