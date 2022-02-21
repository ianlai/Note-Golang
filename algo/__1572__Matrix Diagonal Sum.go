func diagonalSum(mat [][]int) int {
	res := 0
	for i, row := range mat {
		for j, val := range row {
			if i == j || i+j == len(mat)-1 {
				res += val
			}
		}
	}
	return res
}