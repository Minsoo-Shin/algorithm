package maximum_subarray

import "math"

func getSum(nums []int, s, e int) int {
	sum := 0

	for i := s; i < e; i++ {
		sum += nums[i]
	}
	return sum
}

func maxSubArray(nums []int) int {
	maxSum := -math.MaxFloat64
	len := len(nums)

	for i := 1; i < len+1; i++ {
		for k := 0; k < len-i+1; k++ {
			maxSum = math.Max(maxSum, float64(getSum(nums, k, k+i)))
		}
	}
	return int(maxSum)
}
