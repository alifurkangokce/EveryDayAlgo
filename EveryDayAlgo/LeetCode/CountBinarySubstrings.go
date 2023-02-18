package main

// "011000" -> 0, 11, 000 -> [1, 2, 3]
// [1, 2] -> 011 -> (01)
// [2, 3] -> 11000 -> (10, 1100)
// min(1, 2) + min(2, 3) = 3
func countBinarySubstrings(s string) int {
	cur := 1
	pre := 0
	res := 0
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			cur++
		} else {
			res += minCount(cur, pre)
			pre = cur
			cur = 1
		}
	}
	return res + minCount(cur, pre)
}

func minCount(a, b int) int {
	if a < b {
		return a
	}
	return b
}
