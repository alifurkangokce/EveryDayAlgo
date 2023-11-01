package main

import (
	"bufio"
	"fmt"
	"os"
)

func MakeItZero() {
	r := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(r, &t)
	for c := 0; c < t; c++ {
		var n, a int
		fmt.Fscan(r, &n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &a)
		}
		if n%2 == 0 {
			fmt.Printf("2\n1 %d\n1 %d\n", n, n)
		} else {
			fmt.Printf("4\n1 %d\n1 %d\n%d %d\n%d %d\n", n, n-1, n-1, n, n-1, n)
		}
	}
}
