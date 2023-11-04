package main

import (
	"bufio"
	"fmt"
	"os"
)

func GoodArrays() {
	var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tc, n, a int
	for fmt.Fscan(in, &tc); tc > 0; tc-- {
		fmt.Fscan(in, &n)
		x, y := 0, 0
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a)
			x += a
			if a == 1 {
				y += 2
			} else {
				y += 1
			}
		}
		if n == 1 || y > x {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
		}
	}
}
