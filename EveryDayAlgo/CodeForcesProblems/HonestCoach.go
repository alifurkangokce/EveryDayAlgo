package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func HonestCoach() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		s := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &s[j])
		}
		sort.Ints(s)
		min := 1001
		for j := 1; j < n; j++ {
			if min > s[j]-s[j-1] {
				min = s[j] - s[j-1]
			}
		}
		fmt.Fprintln(out, min)
	}
	out.Flush()
}
