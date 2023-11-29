package main

func findNumbers(nums []int) int {
	var res int
	for _, v := range nums {
		var count int
		for v > 0 {
			count++
			v /= 10
		}
		if count%2 == 0 {
			res++
		}
	}
	return res
}
