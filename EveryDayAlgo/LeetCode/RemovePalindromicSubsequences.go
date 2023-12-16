package main

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	} else if s == reverse(s) {
		return 1
	} else {
		return 2
	}
}
