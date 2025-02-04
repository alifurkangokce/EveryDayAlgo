package main

func maximumWealth(accounts [][]int) int {
	max := 0
	for i := 0; i < len(accounts); i++ {
		total := 0
		for j := 0; j < len(accounts[i]); j++ {
			total += accounts[i][j]
		}
		if total > max {
			max = total
		}
	}
	return max
}
