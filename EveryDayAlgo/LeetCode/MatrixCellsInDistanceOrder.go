package main

import "sort"

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	res := make([][]int, 0, R*C)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			res = append(res, []int{i, j})
		}
	}

	sort.Slice(res, func(i, j int) bool {
		dist1 := abs(res[i][0]-r0) + abs(res[i][1]-c0)
		dist2 := abs(res[j][0]-r0) + abs(res[j][1]-c0)
		return dist1 < dist2
	})

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
