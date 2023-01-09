package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func PlusOneOnTheSubset() {
	in := bufio.NewReader(os.Stdin)
	var t, n, res int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a[j])
		}
		sort.Ints(a)
		for i := len(a) - 1; i > 0; i-- {
			if a[i] != a[i-1] {
				res += a[i] - a[i-1]
			}
		}
		fmt.Println(res)
		res = 0
	}
}
