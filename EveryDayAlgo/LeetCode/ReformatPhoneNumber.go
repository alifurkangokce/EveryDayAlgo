package main

import "strings"

func reformatNumber(number string) string {
	number = strings.ReplaceAll(number, " ", "")
	number = strings.ReplaceAll(number, "-", "")
	var res string
	for len(number) > 1 {
		if len(number) == 4 {
			res += number[0:2] + "-"
			res += number[2:4]
			number = ""
		} else if len(number) >= 3 {
			res += number[0:3] + "-"
			number = number[3:]
		} else {
			res += number[0:]
			number = ""
		}
	}
	return strings.Trim(res, "-")
}
