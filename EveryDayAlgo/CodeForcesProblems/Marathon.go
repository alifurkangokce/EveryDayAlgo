package main

import (
	"bufio"
	"fmt"
	"os"
)

func Marathon() {
	in := bufio.NewReader(os.Stdin)
	var t, a, b, c, d, res int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b, &c, &d)
		if a < b {
			res++
		}
		if a < c {
			res++
		}
		if a < d {
			res++
		}
		fmt.Println(res)
		res = 0
	}
}
