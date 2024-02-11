package main

import (
	"fmt"
	"sort"
)

func Puzzles() {
	var n, m int
	fmt.Scan(&n, &m)
	arr := make([]int, m)
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	sort.Ints(arr)
	min := 1000
	for i := 0; i < len(arr)-n+1; i++ {
		if min > arr[i+n-1]-arr[i] {
			min = arr[i+n-1] - arr[i]
		}
	}
	fmt.Print(min)
}
