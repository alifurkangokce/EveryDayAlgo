package main

import "fmt"

func VasyaTheHipster() {
	var a, b int
	fmt.Scan(&a, &b)
	if a > b {
		fmt.Println(b, (a-b)/2)
		return
	}
	fmt.Println(a, (b-a)/2)
}
