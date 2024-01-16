package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func TwoVessels() {
	in := bufio.NewReader(os.Stdin)
	var t int
	var a, b, c float64
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b, &c)
		if a == b {
			fmt.Println(0)
		} else {
			var abs = math.Abs(a - b)
			if abs <= c {
				fmt.Println(1)
			} else {
				res := (abs / c) / 2
				fmt.Println(math.Ceil(res))
			}
		}

	}
}
