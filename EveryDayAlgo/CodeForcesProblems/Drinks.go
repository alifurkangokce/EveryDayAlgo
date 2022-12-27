package main

import "fmt"

func Drinks() {
	var inp int
	var result float64
	fmt.Scan(&inp)
	for i := 0; i < inp; i++ {
		var f float64
		fmt.Scan(&f)
		result += f
	}
	fmt.Print(result / float64(inp))
}
