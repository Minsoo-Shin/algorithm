package valid_palindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		s      string
		result bool
	}{
		{
			s:      "A man, a plan, a canal: Panama",
			result: true,
		},
		{
			s:      "race a car",
			result: false,
		},
		{
			s:      " ",
			result: true,
		},
		{
			s:      "0P",
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, isPalindrome(tt.s))
		assert.Equal(t, tt.result, isPalindrome2(tt.s))
	}
}
