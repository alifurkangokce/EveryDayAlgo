package main

import (
	"bufio"
	"fmt"
	"os"
)

func OrdinaryNumbers() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T int
	fmt.Fscan(in, &T)
	for t := 0; t < T; t++ {
		var n int
		fmt.Fscan(in, &n)

		m := 1
		var tally int
		for i := 1; i <= n; {
			if i == m*9 {
				m = (m * 10) + 1
				i = m
			} else {
				i += m
			}
			tally++
		}

		fmt.Fprintln(out, tally)
	}
}
