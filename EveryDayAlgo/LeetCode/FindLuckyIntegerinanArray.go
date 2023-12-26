package main

func findLucky(arr []int) int {
	mm := make(map[int]int)
	for _, v := range arr {
		mm[v]++
	}
	var max int = -1
	for k, v := range mm {
		if k == v && k > max {
			max = k
		}
	}
	return max
}
