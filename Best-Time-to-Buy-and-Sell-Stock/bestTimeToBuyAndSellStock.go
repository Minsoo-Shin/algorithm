package Best_Time_to_Buy_and_Sell_Stock

import "math"

func maxProfit(prices []int) int {
	// 순차적으로 탐색하면서 내려가면 최저값을 저장한다.
	// 상승하면 최저값의 차이를 계산하면서 크다면 저장한다.
	// maxDelta, minPrice

	maxDelta := 0
	minPrice := math.MaxInt

	for i := 0; i < len(prices); i++ {
		if prices[i] <= minPrice {
			minPrice = prices[i]
		}
		maxDelta = int(math.Max(float64(maxDelta), float64(prices[i]-minPrice)))
	}
	return maxDelta
}
