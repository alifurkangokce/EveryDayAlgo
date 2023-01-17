package main

func findErrorNums(nums []int) []int {
	arr := make(map[int]int)
	var res []int
	var appended []int
	for i := 0; i <= len(nums)-1; i++ {
		arr[nums[i]]++
	}
	for i := 1; i <= len(nums); i++ {
		if arr[i] > 1 {
			res = append(res, i)
		} else if arr[i] != 1 {
			appended = append(appended, i)
		}
	}
	ret := append(res, appended...)
	return ret
}
