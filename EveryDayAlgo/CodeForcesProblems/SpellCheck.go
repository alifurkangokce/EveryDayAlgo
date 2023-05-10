package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SpellCheck() {
	in := bufio.NewReader(os.Stdin)
	var t, n int
	fmt.Fscan(in, &t)
	var s string
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		fmt.Fscan(in, &s)
		if n > 5 || n < 5 || !strings.Contains(s, "T") || !strings.Contains(s, "i") || !strings.Contains(s, "m") || !strings.Contains(s, "u") || !strings.Contains(s, "r") {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}

}
