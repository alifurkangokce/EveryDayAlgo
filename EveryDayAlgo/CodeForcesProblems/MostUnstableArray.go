package main

import (
	"bufio"
	"fmt"
	"os"
)

func MostUnstableArray() {
	in := bufio.NewReader(os.Stdin)
	var t, a, b int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b)
		if a == 1 {
			fmt.Println(0)
		} else if a == 2 {
			fmt.Println(b)
		} else {
			fmt.Println(2 * b)
		}
	}
}
