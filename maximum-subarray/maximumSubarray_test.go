package maximum_subarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximumSubarray(t *testing.T) {
	tests := []struct {
		nums []int
		exp  int
	}{
		{
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			exp:  6,
		},
		{
			nums: []int{1},
			exp:  1,
		},
	}
	for _, tt := range tests {
		got := maxSubArray(tt.nums)
		assert.Equal(t, tt.exp, got)
	}
}
