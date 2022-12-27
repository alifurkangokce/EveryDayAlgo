package main

import "fmt"

func WrongSubtraction() {
	var inp1, inp2 int
	fmt.Scan(&inp1, &inp2)
	for i := 0; i < inp2; i++ {
		if inp1%10 != 0 {
			inp1--
		} else {
			inp1 /= 10
		}
	}
	fmt.Print(inp1)
}
