package main

func longestPalindrome(s string) int {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	var out int
	var odd bool
	for _, vv := range m {
		out += vv
		if vv%2 == 1 {
			out--
			odd = true
		}
	}
	if odd {
		out++
	}
	return out
}
