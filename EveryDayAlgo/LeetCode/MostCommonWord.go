package main

import "strings"

func mostCommonWord(paragraph string, banned []string) string {
	paragraph = strings.ToLower(paragraph)
	mm := make(map[string]int)
	ww := ""
	for i := 0; i <= len(paragraph)-1; i++ {
		if paragraph[i] >= 'a' && paragraph[i] <= 'z' {
			ww += string(paragraph[i])
		} else {
			mm[ww]++
			ww = ""
		}
	}
	if paragraph[len(paragraph)-1] >= 'a' && paragraph[len(paragraph)-1] <= 'z' {
		mm[ww]++
	}

	minVal := 0
	res := ""

	for _, v := range banned {
		if mm[v] > 0 {
			mm[v] = 0
		}
	}
	for k, mv := range mm {
		if mv > minVal && k != "" {
			minVal = mv
			res = k
		}
	}

	return res
}
