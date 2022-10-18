package main

import "fmt"

func Tram() {
	var n, a, b int
	c, m := 0, 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		c -= a
		c += b
		if c > m {
			m = c
		}
	}
	fmt.Println(m)

}
