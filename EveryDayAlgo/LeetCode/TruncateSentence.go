package main

import "strings"

func truncateSentence(s string, k int) string {
	var res string
	ll := strings.Split(s, " ")
	for i := 0; i < k; i++ {
		res += ll[i] + " "
	}
	return strings.Trim(res, " ")
}
