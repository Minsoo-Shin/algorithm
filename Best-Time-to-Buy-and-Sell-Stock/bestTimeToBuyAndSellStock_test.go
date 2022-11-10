package Best_Time_to_Buy_and_Sell_Stock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		result int
	}{
		{
			prices: []int{1, 2, 3},
			result: 2,
		},
		{
			prices: []int{3, 2, 1},
			result: 0,
		},
		{
			prices: []int{1},
			result: 0,
		},
		{
			prices: []int{1, 3, 2, 100},
			result: 99,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, maxProfit(tt.prices))
	}

}
