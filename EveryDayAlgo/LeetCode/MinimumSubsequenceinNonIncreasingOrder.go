package main

import "sort"

func minSubsequence(nums []int) []int {
	var total, totalSub int
	var res []int
	for _, v := range nums {
		total += v
	}
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		totalSub += nums[i]
		total -= nums[i]
		res = append(res, nums[i])
		if totalSub > total {
			break
		}
	}
	return res
}
