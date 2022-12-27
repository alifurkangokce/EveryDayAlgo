package main

import (
	"fmt"
)

func team() {
	var n int
	fmt.Scan(&n)
	var counter int
	for i := 0; i < n; i++ {
		var op1, op2, op3 int
		fmt.Scan(&op1, &op2, &op3)
		if (op1 + op2 + op3) >= 2 {
			counter++
		}
	}
	fmt.Print(counter)
}
