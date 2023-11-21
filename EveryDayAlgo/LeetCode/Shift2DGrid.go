package main

func shiftGrid(grid [][]int, k int) [][]int {
	n := len(grid)
	m := len(grid[0])

	ans := make([][]int, n)
	for i := 0; i < len(ans); i++ {
		ans[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			x := (k / m) % n
			y := k % m

			ans[x][y] = grid[i][j]
			k++
		}
	}

	return ans

}
