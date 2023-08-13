package main

import (
	"fmt"
	"math"
)

func main() {
	got := solution(2, 3)
	if got != 20 {
		fmt.Println("fail", got)
	}
}

func solution(r1 int, r2 int) int64 {
	var count int64
	for i := 1; i <= r2; i++ {
		maxY := math.Floor(math.Sqrt(float64(r2*r2 - i*i)))
		var minY float64
		if i > r1 {
			minY = 0
		} else {
			minY = math.Ceil(math.Sqrt(float64(r1*r1 - i*i)))
		}

		count += int64(maxY - minY + 1.0)
	}
	return count * 4
}
