package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Dragons() {
	in := bufio.NewReader(os.Stdin)
	var s, t int
	fmt.Fscanln(in, &s, &t)
	a := make([][]int, t)
	for it := 0; it < t; it++ {
		var x, y int
		fmt.Fscanln(in, &x, &y)
		a[it] = []int{x, y}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})
	for _, v := range a {
		x, y := v[0], v[1]
		if s > x {
			s += y
		} else {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
