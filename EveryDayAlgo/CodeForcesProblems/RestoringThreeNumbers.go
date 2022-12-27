package main

import (
	"fmt"
	"sort"
)

func RestoringThreeNumbers() {
	var i1, i2, i3, i4 int
	fmt.Scan(&i1, &i2, &i3, &i4)
	arr := []int{i1, i2, i3, i4}
	sort.Ints(arr)
	maks := arr[len(arr)-1]
	for i := 0; i < len(arr)-1; i++ {
		fmt.Printf("%d ", maks-arr[i])
	}

}
