package main

import (
	"bufio"
	"fmt"
	"os"
)

func CandiesAndTwoSisters() {
	var in = bufio.NewReader(os.Stdin)
	var t, n int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		fmt.Println(n/2 - (1 - n%2))
	}

}
