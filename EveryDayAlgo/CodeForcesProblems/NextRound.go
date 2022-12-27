package main

import (
	"fmt"
)

func nextRound() {
	var n, k, count int
	var arr [100]int
	fmt.Scan(&n)
	fmt.Scan(&k)
	for i := 1; i <= n; i++ {
		var l int
		fmt.Scan(&l)
		arr[i] = l
	}
	for j := 1; j <= n; j++ {
		if arr[j] > 0 && arr[j] >= arr[k] {
			count++
		}
	}
	fmt.Print(count)
}
