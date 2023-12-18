package main

func countNegatives(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	r, c := 0, n-1
	ans := 0
	for r < m && c >= 0 {
		if grid[r][c] >= 0 {
			r++
		} else {
			ans += m - r
			c--
		}
	}
	return ans

}
