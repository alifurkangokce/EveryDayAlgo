package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Increasing() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var T int
	for fmt.Fscan(r, &T); T > 0; T-- {

		var n int
		fmt.Fscan(r, &n)
		a := make([]int, n)

		for i := range a {
			fmt.Fscan(r, &a[i])
		}
		sort.Ints(a)

		res := "YES"
		for i := range a {
			if i == 0 {
				continue
			}
			if a[i] <= a[i-1] {
				res = "NO"
				break
			}
		}

		fmt.Fprintln(w, res)
	}

}
