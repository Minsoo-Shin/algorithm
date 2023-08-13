package two_sum

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestTestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		result []int
	}{
		{
			nums:   []int{2, 3, 4},
			target: 5,
			result: []int{0, 1},
		},
		{
			nums:   []int{1, 2, 3, 4},
			target: 7,
			result: []int{2, 3},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			result: []int{1, 2},
		},
	}
	for _, tt := range tests {
		got1 := twoSum1(tt.nums, tt.target)
		sort.Ints(got1)
		sort.Ints(tt.result)
		assert.Equal(t, tt.result, got1)

		got2 := twoSum2(tt.nums, tt.target)
		sort.Ints(got2)
		sort.Ints(tt.result)
		assert.Equal(t, tt.result, got2)

		got3 := twoSum3(tt.nums, tt.target)
		sort.Ints(got3)
		sort.Ints(tt.result)
		assert.Equal(t, tt.result, got3)
	}
}
