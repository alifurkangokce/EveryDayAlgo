package main

import (
	"bufio"
	"fmt"
	"os"
)

func RestorethePermutationbyMerger() {
	var t, n, p int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		mm := make(map[int]int)
		fmt.Fscan(in, &n)
		for i := 0; i < n*2; i++ {
			fmt.Fscan(in, &p)
			if _, ok := mm[p]; !ok {
				fmt.Printf("%d ", p)
				mm[p]++
			}
		}
		fmt.Println()
	}
}
