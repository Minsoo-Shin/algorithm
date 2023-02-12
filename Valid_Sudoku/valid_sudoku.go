package Valid_Sudoku

import "fmt"

func checkSquare(square [][]byte, sr, sc int) bool {
	checkMap := make(map[byte]interface{}, 0)

	for i, _ := range square {
		for j, _ := range square[0] {
			if _, ok := checkMap[square[i][j]]; ok {
				return false
			}
			checkMap[square[i][j]] = struct{}{}
		}
	}

	return true
}

func isValidSudoku(board [][]byte) bool {
	// 가로
	for i := 0; i < len(board); i++ {
		checkMap := make(map[byte]interface{}, 0)

		for _, w := range board[0] {
			if _, ok := checkMap[w]; ok {
				return false
			}
			checkMap[w] = struct{}{}
		}
	}
	fmt.Println("1")
	// 세로
	for j := 0; j < len(board[0]); j++ {
		checkMap := make(map[byte]interface{}, 0)

		for _, v := range board {
			if _, ok := checkMap[v[j]]; ok {
				return false
			}
			checkMap[v[j]] = struct{}{}
		}
	}
	fmt.Println("2")
	// 네모
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if !checkSquare(board, i*3-3, j*3-3) {
				return false
			}
		}
	}
	fmt.Println("3")
	return true
}
