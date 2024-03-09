package main

import (
	"fmt"
)

func OddDivisor() {
	var testcase, n int
	fmt.Scan(&testcase)
	for ; testcase > 0; testcase-- {
		fmt.Scan(&n)
		if (n & (n - 1)) == 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
