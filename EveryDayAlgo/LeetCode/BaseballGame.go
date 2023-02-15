package main

import "strconv"

func calPoints(operations []string) int {
	var result []int
	for _, v := range operations {
		if v == "C" {
			result = append(result[:len(result)-1], result[len(result):]...)
		} else if v == "D" {
			result = append(result, result[len(result)-1]*2)
		} else if v == "+" {
			ss := 0
			for _, r := range result[len(result)-2:] {
				ss += r
			}
			result = append(result, ss)
		} else {
			n, _ := strconv.Atoi(v)
			result = append(result, n)
		}
	}
	res := 0
	for _, rr := range result {
		res += rr
	}
	return res
}
