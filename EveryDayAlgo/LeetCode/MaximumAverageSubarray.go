package main

import "math"

func findMaxAverage(nums []int, k int) float64 {
	var i int
	max := math.MinInt64
	for len(nums)-k-i >= 0 {
		var res int
		for _, v := range nums[i : k+i] {
			res += v
		}
		if res > max {
			max = res
		}
		i++
	}
	return float64(max) / float64(k)
}
