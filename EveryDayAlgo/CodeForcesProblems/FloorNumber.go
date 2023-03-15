package main

import (
	"bufio"
	"fmt"
	"os"
)

func FloorNumber() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	var p, x int
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &p, &x)
		i := 2
		res := 1
		for i < p {
			res++
			i += x
		}
		fmt.Println(res)
	}
}
