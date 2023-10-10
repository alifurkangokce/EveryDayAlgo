package main

import (
	"bufio"
	"fmt"
	"os"
)

func Football() {
	in := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(in, &s)
	var cnt1, cnt0 int
	for _, v := range s {
		if v == '1' {
			cnt1++
			cnt0 = 0
		} else {
			cnt0++
			cnt1 = 0
		}
		if cnt1 >= 7 || cnt0 >= 7 {
			fmt.Print("YES")
			return
		}
	}
	fmt.Print("NO")

}
