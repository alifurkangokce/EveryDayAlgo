package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func Sale() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	var number int
	var arr []int
	fmt.Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &number)
		arr = append(arr, number)
	}
	sort.Ints(arr)
	var res int
	for i := 0; i < m; i++ {
		if arr[i] <= 0 {
			res += arr[i]
		} else {
			break
		}
	}
	fmt.Println(math.Abs(float64(res)))
}
