package main

import (
	"bufio"
	"fmt"
	"os"
)

func YesorYes() {
	in := bufio.NewReader(os.Stdin)
	var t int
	var n string
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		if n == "YES" || n == "yes" || n == "Yes" || n == "YEs" || n == "yES" || n == "yEs" || n == "yeS" || n == "YeS" {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
