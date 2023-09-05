package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var bookCnt, boxWeightLimit int
	fmt.Fscanln(reader, &bookCnt, &boxWeightLimit)

	bookWgts := make([]int, bookCnt)
	for i := 0; i < bookCnt; i++ {
		fmt.Fscan(reader, &bookWgts[i])
	}

	fmt.Println(solution(bookWgts, boxWeightLimit))

	return
}

func solution(wgts []int, boxWeightLimit int) int {
	var (
		boxCount   int = 1
		currBoxWgt int
	)
	if len(wgts) == 0 {
		return 0
	}

	for _, v := range wgts {
		if currBoxWgt+v > boxWeightLimit {
			boxCount++
			currBoxWgt = 0
		}
		currBoxWgt += v
	}
	return boxCount
}
