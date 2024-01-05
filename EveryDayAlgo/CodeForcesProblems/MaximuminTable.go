package main

import (
	"fmt"
)

func MaximumInTable() {
	var n int64
	fmt.Scanln(&n)
	a, b := int64(1), int64(1)
	for i := int64(1); i < n; i++ {
		a *= (n + i - 1)
		b *= i
	}
	fmt.Println(a / b)
}
