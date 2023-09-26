package main

import "sort"

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	var max int = -1
	for i := 0; i < len(nums)-2; i++ {
		if nums[i]+nums[i+1] > nums[i+2] {
			max = i
		}
	}
	if max == -1 {
		return 0
	}
	res := nums[max] + nums[max+1] + nums[max+2]
	return res
}
