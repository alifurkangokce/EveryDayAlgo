package main

import "fmt"

func CoosingTeams() {
	var n, k, res, p int
	fmt.Scan(&n, &k)
	for i := 0; i < n; i++ {
		fmt.Scan(&p)
		if p+k <= 5 {
			res++
		}
	}
	fmt.Print(res / 3)
}
