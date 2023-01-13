package main

import "fmt"

func BlackSquare() {
	var a1, a2, a3, a4, res int
	fmt.Scan(&a1, &a2, &a3, &a4)
	var s string
	fmt.Scan(&s)
	for i := 0; i <= len(s)-1; i++ {
		if s[i] == '1' {
			res += a1
		} else if s[i] == '2' {
			res += a2
		} else if s[i] == '3' {
			res += a3
		} else {
			res += a4
		}
	}
	fmt.Print(res)
}
