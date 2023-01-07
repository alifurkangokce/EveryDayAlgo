package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func TwoArraysAndSwaps() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var t int
	fmt.Fscan(r, &t)

	for i := 0; i < t; i++ {
		var n, k int
		fmt.Fscan(r, &n, &k)
		a := make([]int, n)
		b := make([]int, n)
		s := 0
		for j := 0; j < n; j++ {
			fmt.Fscan(r, &a[j])
			s += a[j]
		}
		for j := 0; j < n; j++ {
			fmt.Fscan(r, &b[j])
		}

		sort.Ints(a)
		sort.Ints(b)
		for j := 0; j < k; j++ {
			if a[j] < b[n-j-1] {
				s += b[n-j-1] - a[j]
			}
		}
		fmt.Fprintln(w, s)
	}
}
