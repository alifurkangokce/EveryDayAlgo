package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func ChipsOnTheBoard() {
	in := bufio.NewReader(os.Stdin)
	var t, n, s int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		var l1, l2 []int
		for k := 0; k < n; k++ {
			fmt.Fscan(in, &s)
			l1 = append(l1, s)
		}
		for k := 0; k < n; k++ {
			fmt.Fscan(in, &s)
			l2 = append(l2, s)
		}
		sort.Ints(l1)
		sort.Ints(l2)
		suml1 := 0
		suml2 := 0
		for _, v := range l1 {
			suml1 += v
		}
		for _, v := range l2 {
			suml2 += v
		}
		if (l1[0]*n)+suml2 <= (l2[0]*n)+suml1 {
			fmt.Println((l1[0] * n) + suml2)
		} else {
			fmt.Println((l2[0] * n) + suml1)
		}

	}
}
