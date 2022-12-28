package main

import "fmt"

func AgainTwentyFive() {
	var n int
	fmt.Scan(&n)
	if n == 0 {
		fmt.Print(0)
	} else if n == 1 {
		fmt.Print(5)
	} else {
		fmt.Print(25)
	}
}
