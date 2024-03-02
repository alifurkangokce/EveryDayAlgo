package main

import (
	"bufio"
	"fmt"
	"os"
)

func MakeitWhite() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, count int
	var s string
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &m, &s)
		count = m
		for j := 0; j < m; j++ {
			if s[j] == 'B' {
				break
			}
			count--
		}
		for j := m - 1; j > 0; j-- {
			if s[j] == 'B' {
				break
			}
			count--
		}
		fmt.Fprintln(out, count)
	}
}
