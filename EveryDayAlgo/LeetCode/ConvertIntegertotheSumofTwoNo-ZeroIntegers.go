package main

import (
	"strconv"
	"strings"
)

func getNoZeroIntegers(n int) []int {
	for A := 1; A < n; A++ {
		B := n - A
		if !strings.Contains(strconv.Itoa(A)+strconv.Itoa(B), "0") {
			return []int{A, B}
		}
	}
	return []int{}
}
