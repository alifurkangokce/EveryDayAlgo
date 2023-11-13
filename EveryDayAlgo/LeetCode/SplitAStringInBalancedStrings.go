package main

func balancedStringSplit(s string) int {
	mm := make(map[rune]int)
	var cnt int
	for _, v := range s {
		mm[v]++
		if mm['R'] == mm['L'] {
			cnt++
			mm['R'] = 0
			mm['L'] = 0
		}
	}
	return cnt
}
