package main

import (
	"fmt"
)

func ArrayWithOddSum() {
	var t int
	for fmt.Scan(&t); t > 0; t-- {
		var n, sum int
		fmt.Scan(&n)
		var odd, even bool
		for i := 0; i < n; i++ {
			var x int
			fmt.Scan(&x)
			sum += x
			if x%2 == 1 {
				odd = true
			} else {
				even = true
			}
		}
		if sum%2 == 1 || (odd && even) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
