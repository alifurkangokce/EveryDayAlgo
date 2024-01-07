package main

func maxPower(s string) int {
	var max, cnt int
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			cnt++
		} else {
			if cnt > max {
				max = cnt
			}
			cnt = 0
		}
	}
	if cnt > max {
		max = cnt
	}
	return max + 1
}
