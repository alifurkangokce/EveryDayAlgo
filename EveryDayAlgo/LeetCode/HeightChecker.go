package main

import "sort"

func heightChecker(heights []int) int {
	var newArr []int
	newArr = append(newArr, heights...)
	sort.Ints(heights)
	var res int
	for i := 0; i <= len(heights)-1; i++ {
		if heights[i] != newArr[i] {
			res++
		}
	}

	return res
}
