package main

func minCostToMoveChips(position []int) int {
	var odd, even int
	for i := 0; i <= len(position)-1; i++ {
		if position[i]%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if even >= odd {
		return odd
	}
	return even
}
