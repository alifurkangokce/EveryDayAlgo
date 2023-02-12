package main

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	max, counter := 0, 1
	for i := 1; i <= len(nums)-1; i++ {
		if nums[i] > nums[i-1] {
			counter++
		} else {
			if counter > max {
				max = counter
			}
			counter = 1
		}
	}
	if counter > max {
		max = counter
	}
	return max

}
