package main

func createTargetArray(nums []int, index []int) []int {
	var res []int
	for i := 0; i <= len(nums)-1; i++ {
		res = append(res[:index[i]], append([]int{nums[i]}, res[index[i]:]...)...)
	}
	return res
}
