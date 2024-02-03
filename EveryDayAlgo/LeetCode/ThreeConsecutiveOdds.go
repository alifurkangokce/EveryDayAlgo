package main

func threeConsecutiveOdds(arr []int) bool {
	var cnt int
	for _, v := range arr {
		if v%2 == 1 {
			cnt++
		} else {
			cnt = 0
		}
		if cnt == 3 {
			return true
		}
	}
	return false
}
