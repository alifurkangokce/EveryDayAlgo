package main

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	sums := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sums[i+1] = sums[i] + nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if sums[i] == sums[len(sums)-1]-sums[i+1] {
			return i
		}
	}
	return -1
}
