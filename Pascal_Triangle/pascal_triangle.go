package Pascal_Triangle

func generate(numRows int) [][]int {
	ret := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		ret[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if i == 0 || j == 0 || j == i {
				ret[i][j] = 1
				continue
			}
			ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
		}
	}
	return ret
}
