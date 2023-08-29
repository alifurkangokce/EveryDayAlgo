package main

import (
	"bufio"
	"fmt"
	"os"
)

func OddSet() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var n, curr, m int
		fmt.Fscan(in, &n)
		for j := 0; j < 2*n; j++ {
			fmt.Fscan(in, &curr)
			if curr%2 == 0 {
				m++
			} else {
				m--
			}
		}
		if m == 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
