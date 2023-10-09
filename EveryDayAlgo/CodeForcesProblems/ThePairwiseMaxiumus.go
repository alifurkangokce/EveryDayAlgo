package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func ThePairwiseMaxiumus() {
	var t int
	scr := bufio.NewReader(os.Stdin)
	ocr := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(scr, &t)

	for i := 0; i < t; i++ {
		a := make([]int, 3)
		fmt.Fscanln(scr, &a[0], &a[1], &a[2])

		sort.Ints(a)
		if a[2] != a[1] {
			fmt.Fprintln(ocr, "NO")
		} else {
			fmt.Fprintln(ocr, "YES")
			fmt.Fprintln(ocr, a[2], 1, a[0])
		}
		ocr.Flush()
	}
}
