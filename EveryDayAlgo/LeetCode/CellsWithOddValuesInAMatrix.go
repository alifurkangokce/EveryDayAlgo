package main

func oddCells(m int, n int, indices [][]int) int {
	var res int
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for i := 0; i <= len(indices)-1; i++ {
		x := indices[i][0]
		y := indices[i][1]
		for k := 0; k <= len(matrix[x])-1; k++ {
			matrix[x][k]++
		}
		for k := 0; k <= len(matrix)-1; k++ {
			matrix[k][y]++
		}
	}
	for _, v := range matrix {
		for _, v2 := range v {
			if v2%2 == 1 {
				res++
			}
		}
	}
	return res
}
