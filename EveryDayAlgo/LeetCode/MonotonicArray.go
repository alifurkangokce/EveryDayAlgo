package main

func isMonotonic(nums []int) bool {
	var mono int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			continue
		} else if nums[i] > nums[i+1] {
			if mono == -1 {
				return false
			}
			mono = 1
		} else {
			if mono == 1 {
				return false
			}
			mono = -1
		}
	}
	return true
}
