package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://www.acmicpc.net/problem/1316
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanln(reader, &t)

	var count int
	for i := 0; i < t; i++ {
		var word string
		fmt.Fscanln(reader, &word)
		if IsGroupWord(word) {
			count++
		}
	}
	fmt.Println(count)
}

func IsGroupWord(w string) bool {
	var checker = make(map[string]bool)

	for i := range w {
		cur := string(w[i])
		_, ok := checker[cur]
		if !ok {
			checker[cur] = true
			continue
		}

		if cur != string(w[i-1]) {
			return false
		}
	}
	return true
}
