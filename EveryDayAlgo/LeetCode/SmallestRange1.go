package main

import "sort"

func smallestRangeI(nums []int, k int) int {
	if len(nums) == 1 {
		return 0
	}
	sort.Ints(nums)
	if (nums[len(nums)-1] - nums[0]) <= (2 * k) {
		return 0
	}
	return (nums[len(nums)-1] - nums[0]) - (2 * k)

}
