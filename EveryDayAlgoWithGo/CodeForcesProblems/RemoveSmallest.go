package main

import (
	"fmt"
	"sort"
)

func RemoveSmallest() {
	var t int
	fmt.Scan(&t)
	for ; t != 0; t-- {
		var n int
		fmt.Scan(&n)
		A := []int{}
		var a int
		for i := 1; i <= n; i++ {
			fmt.Scan(&a)
			A = append(A, a)
		}
		sort.Ints(A)
		ok := true
		for i, j := range A {
			if i > 0 && j-A[i-1] > 1 {
				ok = false
			}
		}
		if ok {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}
