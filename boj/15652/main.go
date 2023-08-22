package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)
	solution(1, n, m, "")
}

func solution(cur, n, m int, temp string) {
	if m == 0 {
		fmt.Println(temp)
		return
	}

	for i := cur; i <= n; i++ {
		solution(i, n, m-1, fmt.Sprintf(temp+"%d ", i))
	}
}
