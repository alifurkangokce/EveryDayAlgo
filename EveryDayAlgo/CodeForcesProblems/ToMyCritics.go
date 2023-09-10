package main

import (
	"bufio"
	"fmt"
	"os"
)

func ToMyCritics() {
	in := bufio.NewReader(os.Stdin)
	var t, a, b, c int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b, &c)
		if a+b >= 10 || a+c >= 10 || c+b >= 10 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
