package main

import "fmt"

func CalculatingFunction() {
	var in int
	fmt.Scan(&in)
	if in%2 == 0 {
		fmt.Print(in / 2)
		return
	}
	fmt.Print(-(in / 2) - 1)
}
