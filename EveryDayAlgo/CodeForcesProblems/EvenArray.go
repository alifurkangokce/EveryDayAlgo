package main

import (
	"bufio"
	"fmt"
	"os"
)

func EvenArray() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		evens, odds := 0, 0
		for j := 0; j < n; j++ {
			var x int
			fmt.Fscan(in, &x)
			if j%2 == 0 && x%2 != 0 {
				evens++
			} else if j%2 != 0 && x%2 == 0 {
				odds++
			}
		}
		if evens == odds {
			fmt.Println(evens)
		} else {
			fmt.Println(-1)
		}
	}
}
