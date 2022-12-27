package main

import "sort"

func thirdMax(nums []int) int {
	sort.Ints(nums)
	cnt := 1
	max := nums[len(nums)-1]
	for i := (len(nums) - 2); i >= 0; i-- {
		if max != nums[i] {
			cnt++
			max = nums[i]
		}
		if cnt == 3 {
			return max
		}
	}
	return nums[len(nums)-1]

}
