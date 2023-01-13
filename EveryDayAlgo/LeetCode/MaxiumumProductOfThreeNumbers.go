package main

import "sort"

func maximumProduct(nums []int) int {
	if len(nums) == 3 {
		return nums[0] * nums[1] * nums[2]
	}
	sort.Ints(nums)
	if nums[len(nums)-1] == 0 {
		return 0
	}
	if nums[len(nums)-1] > 0 && nums[len(nums)-2] > 0 && nums[len(nums)-3] > 0 {
		if nums[0]*nums[1]*nums[len(nums)-1] > nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3] {
			return nums[0] * nums[1] * nums[len(nums)-1]
		} else {
			return nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
		}
	} else if nums[len(nums)-1] < 0 {
		return nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
	} else {
		return nums[0] * nums[1] * nums[len(nums)-1]
	}
}
