package main

import (
	"bufio"
	"fmt"
	"os"
)

func ArrayColoring() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var n, number, sum int
		fmt.Fscan(in, &n)
		var arr []int
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &number)
			arr = append(arr, number)
		}
		for _, v := range arr {
			sum += v
		}
		if sum%2 == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
