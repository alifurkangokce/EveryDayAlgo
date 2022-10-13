package main

import "fmt"

func StonesOnTheTable() {
	var inp int
	var sinp string
	fmt.Scan(&inp, &sinp)
	count := 0
	for i := 0; i < inp-1; i++ {
		if sinp[i] == sinp[i+1] {
			count++
		}
	}
	fmt.Print(count)
}
