package main

import (
	"bufio"
	"fmt"
	"os"
)

func SpecialPermutation() {
	in := bufio.NewReader(os.Stdin)
	var t, n int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		fmt.Printf("%v ", n)
		for i := 1; i < n; i++ {
			fmt.Printf("%v ", i)
		}
		fmt.Println()
	}
}
