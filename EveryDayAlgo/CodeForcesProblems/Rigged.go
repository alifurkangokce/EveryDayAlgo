package main

import (
	"bufio"
	"fmt"
	"os"
)

func Rigged() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var n int
		var ans_w int
		var ans_e int
		f := true
		fmt.Fscan(in, &n, &ans_w, &ans_e)
		for i := 0; i < n-1; i++ {
			var w int
			var e int
			fmt.Fscan(in, &w, &e)
			if w >= ans_w {
				if e >= ans_e {
					f = false
				}
			}
		}
		if f {
			fmt.Println(ans_w)
		} else {
			fmt.Println(-1)
		}
	}
}
