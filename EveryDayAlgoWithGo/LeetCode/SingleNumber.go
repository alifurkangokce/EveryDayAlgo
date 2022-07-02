package main

func singleNumber(nums []int) int {
	r := 0
	for _, n := range nums {
		r = r ^ n
	}
	return r
}
