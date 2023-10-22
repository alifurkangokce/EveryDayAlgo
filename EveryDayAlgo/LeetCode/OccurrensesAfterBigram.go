package main

import "strings"

func findOcurrences(text string, first string, second string) []string {
	arr := strings.Split(text, " ")
	var res []string
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == first && arr[i+1] == second {
			if i+2 <= len(arr)-1 {
				res = append(res, arr[i+2])
			}
		}
	}
	return res
}
