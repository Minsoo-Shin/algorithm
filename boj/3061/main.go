package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var (
		t   int
		ans []int
	)
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanln(reader, &n)

		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &arr[i])
		}

		ans = append(ans, getSwitchCount(arr))
	}

	for i := 0; i < len(ans); i++ {
		fmt.Printf("%v", ans[i])
		if i != len(ans)-1 {
			fmt.Println()
		}
	}

}

func getSwitchCount(arr []int) int {
	var count int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				count++
			}
		}
	}
	return count
}
