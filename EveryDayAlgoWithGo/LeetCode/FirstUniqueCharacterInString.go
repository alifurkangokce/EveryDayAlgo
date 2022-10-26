package main

func firstUniqChar(s string) int {
	x := make(map[byte]int)
	for i := 0; i <= len(s)-1; i++ {
		x[s[i]]++
	}
	for i := 0; i <= len(s)-1; i++ {
		if x[s[i]] == 1 {
			return i
		}
	}
	return -1
}
