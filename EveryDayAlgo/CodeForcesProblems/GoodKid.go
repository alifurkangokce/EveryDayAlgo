package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func GoodKid() {
	var t, n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var arr []int
		var number, res int = 0, 1
		fmt.Fscan(in, &n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &number)
			arr = append(arr, number)
		}
		sort.Ints(arr)
		arr[0]++
		for _, v := range arr {
			res *= v
		}
		fmt.Println(res)
	}
}
