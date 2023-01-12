package main

import "fmt"

func BoringApartments() {
	var t, n, res int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		res = (n % 10) * 10
		if n/1000 >= 1 {
			fmt.Println(res)
			res = 0
		} else if n/100 >= 1 {
			fmt.Println(res - 4)
			res = 0
		} else if n/10 >= 1 {
			fmt.Println(res - 7)
			res = 0
		} else {
			fmt.Println(res - 9)
			res = 0
		}
	}
}
