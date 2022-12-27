package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	i := 0
	var nm1, nm2, el int
	var result, res string
	for i < len(num1) || i < len(num2) {

		if i < len(num1) {
			f, _ := strconv.Atoi(string(num1[len(num1)-1-i]))
			nm1 = f

		}
		if i < len(num2) {

			f, _ := strconv.Atoi(string(num2[len(num2)-1-i]))
			nm2 = f

		}
		if el+nm1+nm2 >= 10 {
			result += fmt.Sprint((nm1 + nm2 + el) - 10)
			el = 1
			nm1 = 0
			nm2 = 0
		} else {
			result += fmt.Sprint((nm1 + nm2 + el))
			el = 0
			nm1 = 0
			nm2 = 0
		}
		i++
	}
	if el == 1 {
		result += "1"
	}
	for i := len(result) - 1; i >= 0; i-- {
		res += string(result[i])
	}

	return string(res)
}
