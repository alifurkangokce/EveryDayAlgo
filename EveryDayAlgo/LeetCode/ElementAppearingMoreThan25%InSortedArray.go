package main

func findSpecialInteger(arr []int) int {
	mm := make(map[int]int)
	for _, v := range arr {
		mm[v]++
	}
	for k, v := range mm {
		if v > len(arr)/4 {
			return k
		}
	}
	return 0
}
