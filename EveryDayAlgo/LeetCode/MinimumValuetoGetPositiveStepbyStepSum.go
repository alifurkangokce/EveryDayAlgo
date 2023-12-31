package main

func minStartValue(nums []int) int {
	if len(nums) == 1 && nums[0] < 0 {
		return absminStartValue(nums[0])
	} else if len(nums) == 1 {
		return 1
	}
	min := 1
	sum := 0
	for i := 0; i <= len(nums)-1; i++ {
		sum += nums[i]
		if nums[i] <= 0 && absminStartValue(sum) >= min && sum < 0 {
			min = absminStartValue(sum) + 1
		}
	}
	return min
}
func absminStartValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
