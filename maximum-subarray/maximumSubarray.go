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

func maxSubArrayN2(nums []int) int {
	maxSum := -math.MaxFloat64
	len := len(nums)

	for i := 0; i < len; i++ {
		sum := 0
		for j := i; j < len; j++ {
			sum += nums[j]
			maxSum = math.Max(maxSum, float64(sum))
		}
	}
	return int(maxSum)
}

func maxSubArrayN1(nums []int) int {
	maxSum, maxSumEnding := nums[0], nums[0]
	len := len(nums)

	for i := 1; i < len; i++ {
		if maxSumEnding < 0 {
			maxSumEnding = nums[i]
		} else {
			maxSumEnding += nums[i]
		}

		maxSum = int(math.Max(float64(maxSum), float64(maxSumEnding)))
	}
	return maxSum
}
