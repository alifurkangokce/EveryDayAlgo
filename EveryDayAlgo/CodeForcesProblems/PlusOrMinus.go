package main

import (
	"bufio"
	"fmt"
	"os"
)

func PlusOrMinus() {
	in := bufio.NewReader(os.Stdin)
	var t, a, b, c int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b, &c)
		if a-b == c {
			fmt.Println("-")
		} else {
			fmt.Println("+")
		}
	}
}
