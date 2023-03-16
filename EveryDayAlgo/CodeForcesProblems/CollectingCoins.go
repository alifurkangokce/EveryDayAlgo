package main

import (
	"bufio"
	"fmt"
	"os"
)

func CollectingCoins() {
	in := bufio.NewReader(os.Stdin)
	var t, a, b, c, n int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b, &c, &n)
		l := (a+b+c+n)%3 == 0
		if !l {
			fmt.Println("NO")
			continue
		}
		x := (a + b + c + n) / 3
		if x < a || x < b || x < c {
			fmt.Println("NO")
			continue
		}
		fmt.Println("YES")
	}
}
