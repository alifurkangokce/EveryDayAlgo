package main

import (
	"bufio"
	"fmt"
	"os"
)

func OddOneOut() {
	var t, a, b, c int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b, &c)
		if a != b && a != c {
			fmt.Println(a)
		} else if b != a && b != c {
			fmt.Println(b)
		} else {
			fmt.Println(c)
		}
	}

}
