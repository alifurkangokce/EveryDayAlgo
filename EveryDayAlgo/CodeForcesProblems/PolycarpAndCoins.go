package main

import (
	"bufio"
	"fmt"
	"os"
)

func PolycarpAndCoins() {
	in := bufio.NewReader(os.Stdin)
	var t, n int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		x := n / 3
		if n%3 == 0 {
			fmt.Printf("%d %d", x, x)
			fmt.Println()
		} else if n%3 == 1 {
			fmt.Printf("%d %d", x+1, x)
			fmt.Println()
		} else {
			fmt.Printf("%d %d", x, x+1)
			fmt.Println()
		}
	}
}
