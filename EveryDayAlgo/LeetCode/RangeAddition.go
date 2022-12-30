package main

func maxCount(m int, n int, ops [][]int) int {
	area := []int{m, n}
	for _, op := range ops {
		intersectionHelper(area, op)
	}
	return area[0] * area[1]
}

func intersectionHelper(a, b []int) {
	a[0] = min(a[0], b[0])
	a[1] = min(a[1], b[1])
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
