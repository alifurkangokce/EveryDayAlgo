package main

import "fmt"

func PrependandAppend() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		var s string
		fmt.Scan(&n, &s)

		l, r := 0, len(s)-1
		for l < r && s[l] != s[r] {
			l++
			r--
		}
		fmt.Println(r - l + 1)
	}
}
