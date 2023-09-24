package main

func repeatedNTimes(nums []int) int {
	mm := make(map[int]int)
	for _, v := range nums {
		mm[v]++
	}
	for k, v := range mm {
		if v*2 == len(nums) && v+1 == len(mm) {
			return k
		}
	}
	return 0
}
