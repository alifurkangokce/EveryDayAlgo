package main

func sumOfUnique(nums []int) int {
	var res int = 0
	mm := make(map[int]int)
	for _, v := range nums {
		mm[v]++
	}
	for k, v := range mm {
		if v == 1 {
			res += k
		}
	}
	return res
}
