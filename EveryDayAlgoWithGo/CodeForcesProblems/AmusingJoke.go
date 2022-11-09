package main

import (
	"fmt"
)

func AmusingJoke() {
	var a, b, c string
	fmt.Scan(&a, &b, &c)
	if len(a)+len(b) != len(c) {
		fmt.Print("NO")
		return
	}
	g := new([256]int)
	r := new([256]int)
	for i := 0; i < len(a); i++ {
		g[a[i]]++
	}
	for i := 0; i < len(b); i++ {
		g[b[i]]++
	}
	for i := 0; i < len(c); i++ {
		r[c[i]]++
	}
	for i := 0; i < 256; i++ {
		if g[i] != r[i] {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
