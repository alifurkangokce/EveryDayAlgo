package main

import "fmt"

func Sum() {
	var t, a, b, c int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&a, &b, &c)
		if a+b == c || a+c == b || b+c == a {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
