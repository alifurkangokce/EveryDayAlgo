package main

func diagonalSum(mat [][]int) int {
	n, res := len(mat), 0

	for i := 0; i < n; i++ {
		res += mat[i][i]
		if i != n-i-1 {
			res += mat[i][n-i-1]
		}
	}
	return res

}
