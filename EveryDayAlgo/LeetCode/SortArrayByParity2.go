package main

func sortArrayByParityII(nums []int) []int {
	var arrOdd, arrEven, res []int
	for i := 0; i <= len(nums)-1; i++ {
		if nums[i]%2 == 0 {
			arrEven = append(arrEven, nums[i])
		} else {
			arrOdd = append(arrOdd, nums[i])
		}
	}
	for i := 0; i <= len(arrEven)-1; i++ {
		res = append(res, arrEven[i])
		res = append(res, arrOdd[i])
	}
	return res
}
