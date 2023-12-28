package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func LeastProduct() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	n := 1
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var n int
		fmt.Fscan(in, &n)
		arr := make([]int64, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		cur := int64(1)
		for _, v := range arr {
			if v < 0 {
				cur *= -1
			} else if v == 0 {
				cur = 0
				break
			}
		}
		if cur > 0 {
			fmt.Fprintln(out, 1)
			fmt.Fprintf(out, "%d %d\n", 1, 0)
		} else {
			fmt.Fprintln(out, 0)
		}
	}
}
