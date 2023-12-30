package main

import "strings"

func stringMatching(words []string) []string {
	var res []string
	for i := 0; i <= len(words)-1; i++ {
		for j := 0; j <= len(words)-1; j++ {
			if j != i {
				if strings.Contains(words[j], words[i]) {
					res = append(res, words[i])
					break
				}
			}
		}
	}
	return res
}
