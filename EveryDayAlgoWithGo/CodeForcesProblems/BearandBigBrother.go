package main

import "fmt"

func BearAndBigBrother() {
	var inp1, inp2, count int
	fmt.Scan(&inp1, &inp2)
	for inp1 <= inp2 {
		count++
		inp1 *= 3
		inp2 *= 2
	}
	fmt.Print(count)
}
