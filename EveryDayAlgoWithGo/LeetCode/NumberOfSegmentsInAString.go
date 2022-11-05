package main

func countSegments(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && (i == 0 || s[i-1] == ' ') {
			res++
		}
	}

	return res
}
