package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	dp [][]int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanln(reader, &t)
	dp = make([][]int, 31)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 31)
	}
	for i := 0; i < t; i++ {
		var n, m int
		fmt.Fscanln(reader, &n, &m)
		fmt.Fprintln(writer, Combination2(m, n))
	}
}

/*
Combination 계산하기
*/
//
//func Combination1(m, n int) int {
//	if m == n || n == 0 {
//		return 1
//	}
//	if dp[m][n] != 0 {
//		return dp[m][n]
//	}
//	dp[m][n] = Combination1(m-1, n) + Combination1(m-1, n-1) // 조합의 성질, 파스칼의 삼각형 활용.
//	return dp[m][n]
//}

func Combination2(m, n int) int {
	mul := 1

	for i := 0; i < n; i++ {
		mul = mul * (m - i) / (i + 1)
	}

	return mul
}
