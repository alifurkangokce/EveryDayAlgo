package main

import "fmt"

func Cplusequal() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var a, b, n int
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&n)
		var ans int = 0
		for {
			if a > n || b > n {
				break
			}
			if a < b {
				a += b
			} else {
				b += a
			}
			ans++
		}
		fmt.Println(ans)
	}
}
