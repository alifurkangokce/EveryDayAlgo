package main

import "fmt"

func ILoveUserName() {
	var inp, result, inpp int
	fmt.Scan(&inp, &inpp)
	max, min := inpp, inpp
	for i := 0; i < inp-1; i++ {
		fmt.Scan(&inpp)
		if inpp < min {
			result++
			min = inpp
		}
		if inpp > max {
			result++
			max = inpp
		}
	}
	fmt.Print(result)
}
