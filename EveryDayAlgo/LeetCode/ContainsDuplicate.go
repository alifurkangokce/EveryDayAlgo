package main

func containsDuplicate(nums []int) bool {
	//sort.Ints(nums)
	// x := make(map[int]bool)
	// for _, n := range nums {
	// 	if x[n] {
	// 		return true
	// 	}
	// 	x[n] = true
	// }
	// return false

	// x := nums[0]
	// for i := 1; i < len(nums); i++ {
	// 	if x == nums[i] {
	// 		return true
	// 	}
	// 	x = nums[i]
	// }
	// return false

	m := map[int]int{}

	for _, val := range nums {
		m[val]++
		if m[val] > 1 {
			return true
		}
	}

	return false
}
