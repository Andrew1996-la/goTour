package multiple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiple(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "positive",
			a:        2,
			b:        3,
			expected: 6,
		},
		{
			name:     "multiply by zero",
			a:        10,
			b:        0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)

			assert.Equal(
				t,
				tt.expected,
				result,
			)
		})
	}
}
