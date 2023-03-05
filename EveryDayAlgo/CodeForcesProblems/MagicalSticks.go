package main

import "fmt"

func MagicalSticks() {
	var t int
	var n int
	fmt.Scanln(&t)
	for ; t > 0; t-- {
		fmt.Scanln(&n)
		fmt.Println(((n - 1) / 2) + 1)
	}
}
