package main

func isToeplitzMatrix(matrix [][]int) bool {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return true
	}

	m, n := len(matrix), len(matrix[0])
	x, y := len(matrix)-1, 0
	for x >= 0 && y < n {
		for xx, yy := x, y; xx < m && yy < n; xx, yy = xx+1, yy+1 {
			if matrix[xx][yy] != matrix[x][y] {
				return false
			}
		}

		if x > 0 {
			x--
		} else {
			y++
		}
	}

	return true
}
