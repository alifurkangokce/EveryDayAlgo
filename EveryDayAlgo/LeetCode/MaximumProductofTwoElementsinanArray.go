package main

import "sort"

func maxProduct(nums []int) int {
	sort.Ints(nums)
	ll := len(nums)
	return (nums[ll-1] - 1) * (nums[ll-2] - 1)
}
