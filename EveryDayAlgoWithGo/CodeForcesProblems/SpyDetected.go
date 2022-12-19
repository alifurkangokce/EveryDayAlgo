package main

import (
	"bufio"
	"fmt"
	"os"
)

func SpyDetected() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--
		var n int
		fmt.Fscan(in, &n)
		v := make([]int, n)
		for i := range v {
			fmt.Fscan(in, &v[i])
		}
		for i := range v {
			if v[i] != v[(i+1)%n] && v[i] != v[(i+2)%n] {
				fmt.Println(i + 1)
				break
			}
		}
	}
}
