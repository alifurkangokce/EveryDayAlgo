package main

var directions = [][2]int{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
	{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
}

func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			sum, cells := img[i][j], 1
			for _, dir := range directions {
				x, y := i+dir[0], j+dir[1]
				if x >= 0 && x < m && y >= 0 && y < n {
					sum += img[x][y]
					cells++
				}
			}
			res[i][j] = sum / cells
		}
	}
	return res
}
