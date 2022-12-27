package main

import "fmt"

func VanyaAndFence() {
	var inp1, dis, res int
	fmt.Scan(&inp1, &dis)
	for i := 0; i < inp1; i++ {
		var d int
		fmt.Scan(&d)
		if dis >= d {
			res++
		} else {
			res += 2
		}
	}
	fmt.Print(res)
}
