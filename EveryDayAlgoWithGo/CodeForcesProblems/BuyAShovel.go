package main

import "fmt"

func BuyAShovel() {
	var k, r int
	fmt.Scan(&k, &r)
	for i := 1; true; i++ {
		res := (k * i) % 10
		if res == 0 || res == r {
			fmt.Print(i)
			return
		}
	}
	fmt.Print(0)
}
