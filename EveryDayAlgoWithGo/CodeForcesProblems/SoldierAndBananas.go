package main

import "fmt"

func SoldierAndBananas() {
	var k, n, w, x int
	fmt.Scan(&k, &n, &w)
	for i := 1; i <= w; i++ {
		x += (k * i)
	}
	if x <= n {
		fmt.Print(0)
	} else {
		fmt.Print(x - n)
	}
}
