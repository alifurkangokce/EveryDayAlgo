package main

import (
	"bufio"
	"fmt"
	"os"
)

func LCMProblem() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	var n1, n2 int
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n1, &n2)
		if n1*2 <= n2 {
			fmt.Printf("%d %d\n", n1, n1*2)
		} else {
			fmt.Println("-1 -1")
		}
	}
}
