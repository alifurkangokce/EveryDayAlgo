package main

func decompressRLElist(nums []int) []int {
	var res []int
	for i := 0; i <= len(nums)-1; i += 2 {
		for j := 0; j < nums[i]; j++ {
			res = append(res, nums[i+1])
		}
	}
	return res
}
