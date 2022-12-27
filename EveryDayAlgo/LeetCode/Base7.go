package main

import "strconv"

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	res := ""
	ngCheck := true
	if num < 0 {
		ngCheck = false
		num = num * -1
	}
	for num > 0 {
		res = strconv.Itoa(num%7) + res
		num /= 7
	}
	if ngCheck {
		return res
	} else {
		return "-" + res
	}
}
