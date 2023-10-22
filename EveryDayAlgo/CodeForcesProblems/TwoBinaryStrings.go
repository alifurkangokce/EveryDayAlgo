package main

import (
	"bufio"
	"fmt"
	"os"
)

func TwoBinaryStrings() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var a, b string
		isTrue := false
		fmt.Fscan(in, &a, &b)
		for i := 0; i < len(a)-1; i++ {
			if a[i] == '0' && a[i] == b[i] && a[i+1] == b[i+1] && a[i+1] == '1' {
				fmt.Println("YES")
				isTrue = true
				break
			}
		}
		if !isTrue {
			fmt.Println("NO")
		}

	}

}
