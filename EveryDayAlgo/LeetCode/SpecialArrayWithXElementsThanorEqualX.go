package main

import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)
	if nums[0] >= len(nums) {
		return len(nums)
	}
	for i, n := range nums {
		if n >= len(nums)-i && i > 0 && nums[i-1] < len(nums)-i {
			return len(nums) - i
		}
	}

	return -1
}
