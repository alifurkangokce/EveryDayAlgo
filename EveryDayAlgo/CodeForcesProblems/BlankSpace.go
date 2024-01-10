package main

import (
	"bufio"
	"fmt"
	"os"
)

func BlankSpace() {
	in := bufio.NewReader(os.Stdin)
	var t, n, a int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		var max, cnt int
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a)
			if a == 0 {
				cnt++
			} else {
				if cnt > max {
					max = cnt
				}
				cnt = 0
			}
		}
		if cnt > max {
			max = cnt
		}
		fmt.Println(max)
	}
}
