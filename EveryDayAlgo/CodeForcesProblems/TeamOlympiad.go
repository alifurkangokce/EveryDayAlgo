package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a int, rest ...int) int {
	for _, r := range rest {
		if r < a {
			a = r
		}
	}
	return a
}

func TeamOlympiad() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	m := make([][]int, 3)
	m[0] = []int{}
	m[1] = []int{}
	m[2] = []int{}

	for i := 0; i < n; i++ {
		var v int
		fmt.Fscan(in, &v)

		switch v {
		case 1:
			m[0] = append(m[0], i+1)
		case 2:
			m[1] = append(m[1], i+1)
		case 3:
			m[2] = append(m[2], i+1)
		}
	}

	l := min(len(m[0]), len(m[1]), len(m[2]))
	fmt.Fprintln(out, l)
	for i := 0; i < l; i++ {
		fmt.Fprintf(out, "%d %d %d\n", m[0][i], m[1][i], m[2][i])
	}
}
