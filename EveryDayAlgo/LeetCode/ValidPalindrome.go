package main

import "strings"

func isPalindromeString(s string) bool {
	t2 := []byte{}
	s = strings.ToLower(s)
	for i := 0; i < len(s); i++ {
		if (s[i] >= '0' && s[i] <= '9') || s[i]-'a' <= 25 {
			t2 = append(t2, s[i])
		}

	}

	for i := 0; i < len(t2)/2; i++ {
		if t2[i] != t2[len(t2)-1-i] {
			return false
		}
	}
	return true
}
