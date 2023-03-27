package main

import (
	"bufio"
	"fmt"
	"os"
)

func CodeForcesChecking() {
	mm := make(map[string]bool)
	mm["c"] = true
	mm["o"] = true
	mm["d"] = true
	mm["e"] = true
	mm["f"] = true
	mm["r"] = true
	mm["s"] = true
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	var inp string
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &inp)
		if mm[inp] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
