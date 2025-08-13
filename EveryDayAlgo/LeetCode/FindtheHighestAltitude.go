package main

func largestAltitude(gain []int) int {
	var res, val int = 0, gain[0]
	if gain[0] > 0 {
		res = gain[0]
	}
	for i := 1; i <= len(gain)-1; i++ {
		if gain[i]+val > res {
			res = gain[i] + val
		}
		val = gain[i] + val
	}
	return res
}
