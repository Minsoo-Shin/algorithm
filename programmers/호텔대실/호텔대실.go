package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

// [시작, 끝]
//

// O(N^2)
// 일단 끝나는 시간 기준으로 정렬한다.
// 끝나는 시간 + 10분 보다 다른 시작시간이 더 앞에 있는 룸의 갯수를 최고 갯수와 비교하며
// 최고 값을 찾아간다.

func transferMinutes(clock string) int {
	time := strings.Split(clock, ":")

	hour, err := strconv.Atoi(time[0])
	if err != nil {
		return 0
	}

	minutes, err := strconv.Atoi(time[1])
	if err != nil {
		return 0
	}

	return 60*hour + minutes
}

func 둘만의암호(bookTime [][]string) int {
	var a [][]int
	for _, reserv := range bookTime {
		startMinutes := transferMinutes(reserv[0])
		endMinutes := transferMinutes(reserv[1])

		a = append(a, []int{startMinutes, endMinutes})
	}

	sort.Slice(a, func(i, j int) bool {
		if a[i][1] <= a[j][1] {
			return true
		}
		return false
	})

	maxInt := 0
	for i := 0; i < len(a); i++ {
		overlapInt := 1
		for j := i + 1; j < len(a); j++ {
			if i == j {
				continue
			}
			if a[j][0] < a[i][1]+10 {
				overlapInt++
			}
		}
		maxInt = int(math.Max(float64(maxInt), float64(overlapInt)))
	}
	return maxInt
}
