package main

import "fmt"

func VanyaAndCubes() {
	var n int
	fmt.Scanf("%d", &n)
	sum := 0
	for i := 1; true; i++ {
		sum += i
		if sum > n {
			fmt.Println(i - 1)
			return
		}
		n -= sum
	}
}
