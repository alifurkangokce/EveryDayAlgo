package main

import "fmt"

func HitTheLottery() {
	var n, k int
	fmt.Scan(&n)
	for _, i := range []int{100, 20, 10, 5, 1} {
		k += n / i
		n %= i
	}
	fmt.Println(k)
}
