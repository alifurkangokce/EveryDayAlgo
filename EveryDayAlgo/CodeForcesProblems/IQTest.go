package main

import (
	"bufio"
	"fmt"
	"os"
)

func IqTest() {
	in := bufio.NewReader(os.Stdin)
	var t, n, dif1, dif2, r1, r2 int
	var arr []int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		arr = append(arr, n)
	}
	for k, v := range arr {
		if v%2 == 0 {
			dif1++
			r1 = k
		} else {
			dif2++
			r2 = k
		}
		if dif1 > dif2 && dif2 != 0 {
			fmt.Print(r2 + 1)
			return
		} else if dif2 > dif1 && dif1 != 0 {
			fmt.Print(r1 + 1)
			return
		}
	}
}
