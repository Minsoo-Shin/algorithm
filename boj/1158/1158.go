package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://www.acmicpc.net/problem/1158
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &m, &n)
	fmt.Fprintf(writer, "<%v>\n", strings.Join(solution(m, n), ", "))
}

func solution(n, k int) []string {
	var ret = make([]string, 0)
	var circle = make([]int, n)
	for i := 1; i <= n; i++ {
		circle[i-1] = i
	}
	cur := 0
	for len(circle) != 0 {
		size := len(circle)
		cur = (cur + k - 1) % size
		ret = append(ret, strconv.Itoa(circle[cur]))
		circle = append(circle[:cur], circle[cur+1:]...)
	}
	return ret
}
