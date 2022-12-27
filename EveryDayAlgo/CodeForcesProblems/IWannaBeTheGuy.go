package main

import "fmt"

func IWannaBeTheGuy() {
	var c, lx, ly int
	fmt.Scan(&c)
	x := make(map[int]bool)
	fmt.Scan(&lx)
	for i := 1; i <= c; i++ {
		x[i] = false
	}
	for i := 0; i < lx; i++ {
		var inp1 int
		fmt.Scan(&inp1)
		x[inp1] = true
	}
	fmt.Scan(&ly)
	for i := 0; i < ly; i++ {
		var inp2 int
		fmt.Scan(&inp2)
		x[inp2] = true
	}
	for i := 1; i <= c; i++ {
		if !x[i] {
			fmt.Print("Oh, my keyboard!")
			return
		}
	}
	fmt.Print("I become the guy.")
}
