package main

import (
	"bufio"
	"fmt"
	"os"
)

func ShortSort() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var s string
		fmt.Fscan(in, &s)

		if (s[0] == 'c' && s[1] == 'a') || (s[1] == 'c' && s[2] == 'a') {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}

	}
}
