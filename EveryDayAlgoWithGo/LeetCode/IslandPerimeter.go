package main

func islandPerimeter(grid [][]int) int {

	res := 0
	for i := 0; i < len(grid); i++ {
		pre := 0
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				res += 4
			}
		}
	}
	return 0
}
