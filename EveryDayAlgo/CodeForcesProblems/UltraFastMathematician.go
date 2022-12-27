package main

import (
	"fmt"
)

func UltraFastMathematician() {
	var a, b string
	fmt.Scan(&a, &b)
	for i, val := range a {
		if val == rune(b[i]) {
			fmt.Print("0")
		} else {
			fmt.Print("1")
		}
	}
}
