package main

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	arr := make(map[string]int)
	s1Arr := strings.Split(s1, " ")
	s2Arr := strings.Split(s2, " ")
	var res []string
	for _, v := range s1Arr {
		arr[v]++
	}
	for _, v := range s2Arr {
		arr[v]++
	}
	for k, v := range arr {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
