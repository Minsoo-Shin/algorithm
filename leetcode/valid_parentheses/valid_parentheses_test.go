package valid_parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		s      string
		result bool
	}{
		{
			s:      "{}",
			result: true,
		},
		{
			s:      "{(})",
			result: false,
		},
		{
			s:      "}{",
			result: false,
		},
		{
			s:      "{}()[]",
			result: true,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, isValid(tt.s), tt.result)
	}
}
