package main

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)

	var sum int
	min := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if k > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
		sum += nums[i]
		if nums[i] < min {
			min = nums[i]
		}
	}

	if k%2 != 0 {
		sum -= 2 * min
	}
	return sum

}
