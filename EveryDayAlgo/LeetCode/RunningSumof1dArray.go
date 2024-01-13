package main

func runningSum(nums []int) []int {
	var res []int
	res = append(res, nums[0])
	if len(nums) == 1 {
		return res
	}
	for i := 1; i <= len(nums)-1; i++ {
		res = append(res, res[i-1]+nums[i])
	}
	return res
}
