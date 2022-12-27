package main

import (
	"bufio"
	"fmt"
	"os"
)

func YetAnotherTwoIntegersProblem() {
	var in1 = bufio.NewReader(os.Stdin)
	var t, n1, n2 int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in1, &n1, &n2)
		if (n1-n2)%10 != 0 {
			x := (n1 - n2) / 10
			if x < 0 {
				fmt.Println((x * -1) + 1)
			} else {
				fmt.Println(x + 1)
			}
		} else {
			x := (n1 - n2) / 10
			if x < 0 {
				fmt.Println(x * -1)
			} else {
				fmt.Println(x)
			}
		}
	}
}
