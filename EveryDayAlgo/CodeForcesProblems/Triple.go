package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Triple() {
	in := bufio.NewReader(os.Stdin)
	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}
		sort.Ints(arr)
		ans := -1
		for i := 1; i < n-1; i++ {
			if arr[i] == arr[i-1] && arr[i] == arr[i+1] {
				ans = arr[i]
			}
		}
		fmt.Println(ans)
	}
}
