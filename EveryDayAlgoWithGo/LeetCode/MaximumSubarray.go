package main

func maxSubArray(nums []int) int {
	// if len(nums) == 0 {
	// 	return 0
	// }
	// max := nums[0]
	// result := nums[0]
	// for i := 1; i <= len(nums)-1; i++ {
	// 	result += nums[i]
	// 	if result > max && nums[i] > result {
	// 		max = nums[i]
	// 		result = max
	// 	} else if result > max {
	// 		max = result
	// 		result = max
	// 	} else if nums[i] > max && result > 0 {
	// 		max = nums[i] + result
	// 		result = max
	// 	} else if nums[i] > max && result < 0 {
	// 		max = nums[i]
	// 		result = max
	// 	} else if result < 0 {
	// 		result = 0
	// 	}

	// }
	// return max
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	result := nums[0]
	for i := 1; i <= len(nums)-1; i++ {
		if result > 0 {
			result += nums[i]
		} else {
			result = nums[i]
		}
		if result > max {
			max = result
		}
	}
	return max

}
