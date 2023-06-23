package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeBase62(t *testing.T) {
	test_values := []uint{0, 17, 61, 62, 3844}
	expected_values := []string{"0", "h", "Z", "10", "100"}
	for i, test_value := range test_values {
		expected_value := expected_values[i]
		actual_value := EncodeBase62(test_value)

		assert.Equal(t, expected_value, actual_value)
	}
}
