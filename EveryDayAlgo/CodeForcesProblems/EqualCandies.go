package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func EqualCandies() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	var tt int
	var number int
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &tt)
		var arr []int
		var res int
		for i := 0; i < tt; i++ {
			fmt.Fscan(in, &number)
			arr = append(arr, number)
		}
		sort.Ints(arr)
		for i := len(arr) - 1; i > 0; i-- {
			res += arr[i] - arr[0]
		}
		fmt.Println(res)
	}

}
