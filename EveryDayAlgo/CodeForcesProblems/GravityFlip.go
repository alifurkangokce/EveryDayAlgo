package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func GravityFlip() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)
	mm := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &mm[i])
	}
	sort.Ints(mm)
	for _, v := range mm {
		fmt.Printf("%d ", v)
	}
}
