package main

import (
	"bufio"
	"fmt"
	"os"
)

func EvenOdds() {
	in := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(in, &n, &k)
	if n%2 != 0 {
		n = n + 1
	}
	if n/2 < k {
		fmt.Print(2 * (k - (n / 2)))
	} else {
		fmt.Print(2*k - 1)
	}
}
