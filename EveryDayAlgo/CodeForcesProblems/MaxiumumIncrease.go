package main

import (
	"bufio"
	"fmt"
	"os"
)

func MaxiumumIncrease() {
	in := bufio.NewReader(os.Stdin)
	var t, max int
	res, max := 1, 1
	fmt.Fscan(in, &t)
	var arr []int
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		arr = append(arr, n)
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] <= arr[i+1] {
			if arr[i] != arr[i+1] {
				res++
			} else {
				res = 1
			}
		} else {
			res = 1
		}
		if res > max {
			max = res

		}
	}
	fmt.Print(max)
}
