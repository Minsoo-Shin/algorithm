package valid_anagram

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	tests := []struct {
		s      string
		t      string
		result bool
	}{
		{
			s:      "ana",
			t:      "nna",
			result: false,
		},
		{
			s:      "anagram",
			t:      "naagram",
			result: true,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, isAnagram(tt.s, tt.t))
		assert.Equal(t, tt.result, isAnagram2(tt.s, tt.t))
	}
}
