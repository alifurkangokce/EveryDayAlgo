package main

import "fmt"

func Elephant() {
	var i int
	fmt.Scan(&i)
	if (i % 5) != 0 {
		fmt.Print((i / 5) + 1)
	} else {
		fmt.Print((i / 5))
	}
}
