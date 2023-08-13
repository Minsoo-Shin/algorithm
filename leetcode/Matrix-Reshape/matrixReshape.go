package Matrix_Reshape

func getReshapeSize(m, n, r, c int) (int, int) {
	if m*n == r*c {
		return r, (m * n) / r
	}

	return m, n
}

func matrixReshape(mat [][]int, r int, c int) [][]int {

	rowLen, colLen := getReshapeSize(len(mat), len(mat[0]), r, c)

	ret := make([][]int, rowLen)
	for i, _ := range ret {
		ret[i] = make([]int, colLen)
	}

	count := 0
	for i, _ := range mat {
		for j, _ := range mat[0] {

			k, l := count/colLen, count%colLen

			ret[k][l] = mat[i][j]
			count++
		}
	}

	return ret
}
