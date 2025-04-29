package main

import "strings"

func halvesAreAlike(s string) bool {
	vovels := "aeiouAEIOU"
	s1 := s[:len(s)/2]
	s2 := s[len(s)/2:]
	var cnt1, cnt2 int
	for i := 0; i <= len(s1)-1; i++ {
		if strings.Contains(vovels, string(s1[i])) {
			cnt1++
		}
		if strings.Contains(vovels, string(s2[i])) {
			cnt2++
		}
	}
	return cnt1 == cnt2

}
