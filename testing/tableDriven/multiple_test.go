package multiple

import "testing"

func TestMultiple(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "positive numbers",
			a:        2,
			b:        2,
			expected: 4,
		},
		{
			name:     "negative and positive",
			a:        -2,
			b:        3,
			expected: -6,
		},
		{
			name:     "multiply by zero",
			a:        10,
			b:        0,
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Multiply(test.a, test.b)

			if result != test.expected {
				t.Errorf(
					"expected %d, got %d",
					test.expected,
					result,
				)
			}
		})
	}
}
