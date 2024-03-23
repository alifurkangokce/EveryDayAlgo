package main

func maxDepthNesting(s string) int {
	var max int
	var cnt int
	for i := 0; i <= len(s)-1; i++ {
		if s[i] == '(' {
			cnt++
		}
		if s[i] == ')' {
			cnt--
		}
		if cnt > max {
			max = cnt
		}
	}
	return max
}
