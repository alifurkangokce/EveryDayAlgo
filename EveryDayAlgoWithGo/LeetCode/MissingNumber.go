package main

func missingNumber(nums []int) int {
	var result int
	x := (len(nums) * (len(nums) + 1)) / 2
	for _, v := range nums {
		result += v
	}
	return x - result
}
