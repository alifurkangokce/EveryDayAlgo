package main

import (
	"fmt"
)

func SerejaAndDima() {
	var n int
	score := map[int]int{}
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	first, last, counter := 0, n-1, 0
	for first <= last {
		if arr[first] > arr[last] {
			score[counter%2] += arr[first]
			first++
		} else {
			score[counter%2] += arr[last]
			last--
		}
		counter++
	}

	fmt.Println(score[0], score[1])
}
