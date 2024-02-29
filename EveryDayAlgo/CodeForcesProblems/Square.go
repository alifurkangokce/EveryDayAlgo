package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Square() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	var a, b int
	for i := 0; i < t; i++ {
		var arr1 []int
		for i := 0; i < 4; i++ {
			fmt.Fscan(in, &a, &b)
			arr1 = append(arr1, a)
		}
		sort.Ints(arr1)
		dd := arr1[3] - arr1[0]
		fmt.Println(dd * dd)
	}
}
