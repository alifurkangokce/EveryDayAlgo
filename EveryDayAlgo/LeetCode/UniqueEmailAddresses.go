package main

import "strings"

func numUniqueEmails(emails []string) int {
	res := make(map[string]int)
	for _, v := range emails {
		arr := strings.Split(v, "@")
		var ff string
		for _, v := range arr[0] {
			if v != '+' {
				ff += string(v)
			} else {
				break
			}
		}

		newff := strings.ReplaceAll(ff, ".", "")
		newff += "@" + arr[1]
		res[newff]++
	}
	return len(res)
}
