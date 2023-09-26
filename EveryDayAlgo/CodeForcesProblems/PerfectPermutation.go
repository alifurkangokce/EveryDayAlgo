package main

import (
	"bufio"
	"fmt"
	"os"
)

func PerfectPermutation() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)
	if n%2 != 0 {
		fmt.Print(-1)
	} else {
		for i := 1; i <= n; i += 2 {
			fmt.Printf("%d ", i+1)
			fmt.Printf("%d ", i)
		}
	}
}
