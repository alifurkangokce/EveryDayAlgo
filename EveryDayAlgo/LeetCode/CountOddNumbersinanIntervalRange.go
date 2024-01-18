package main

func countOdds(low int, high int) int {
	var res int
	if low%2 == 0 && high%2 == 0 {
		res = (high - low) / 2
	} else {
		res = (high-low)/2 + 1
	}
	return res
}
