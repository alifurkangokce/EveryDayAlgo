package main

import "fmt"

func BalancedArray() {
	var inp, n int
	fmt.Scan(&inp)
	for i := 0; i < inp; i++ {
		fmt.Scan(&n)
		if n%4 != 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
			for j := 2; j <= n; j += 2 {
				fmt.Printf("%d ", j)
			}
			for j := 1; j < n-1; j += 2 {
				fmt.Printf("%d ", j)
			}
			fmt.Printf("%d ", n+n/2-1)
		}
	}
}
