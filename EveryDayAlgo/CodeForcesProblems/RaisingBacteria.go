package main

import (
	"fmt"
)

func RaisingBacteria() {
	var x int
	fmt.Scan(&x)
	cnt := 0
	for x != 0 {
		cnt += x & 1
		x /= 2
	}
	fmt.Println(cnt)
}
