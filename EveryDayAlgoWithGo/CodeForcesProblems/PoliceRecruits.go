package main

import (
	"bufio"
	"fmt"
	"os"
)

func PoliceRecruits() {
	var in, cnt, v, c int
	inp := bufio.NewReader(os.Stdin)
	fmt.Fscan(inp, &in)
	for i := 0; i < in; i++ {
		fmt.Fscan(inp, &v)
		c += v
		if c < 0 {
			cnt++
			c = 0
		}
	}
	fmt.Print(cnt)
}
