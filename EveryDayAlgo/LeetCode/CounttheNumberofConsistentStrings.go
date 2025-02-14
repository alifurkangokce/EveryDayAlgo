package main

import "strings"

func countConsistentStrings(allowed string, words []string) int {
	var res int
	for _, v := range words {
		check := true
		for i := 0; i <= len(v)-1; i++ {
			if !strings.Contains(allowed, string(v[i])) {
				check = false
				break
			}
		}
		if check {
			res++
		}
	}
	return res

}
