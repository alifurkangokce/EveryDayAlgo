package main

import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
	var n []int
	var res int = 0
	for i := 0; i <= len(points)-1; i++ {
		n = append(n, points[i][0])
	}
	sort.Ints(n)
	for i := 1; i <= len(n)-1; i++ {
		r1 := n[i] - n[i-1]
		if r1 > res {
			res = r1
		}
	}
	return res
}
