package main

import (
	"fmt"
)

func BachgoldProblem() {
	var a int
	fmt.Scan(&a)
	fmt.Println(a / 2)
	if a%2 == 1 {
		fmt.Print("3 ")
	} else {
		fmt.Print("2 ")
	}
	for i := 1; i <= (a-a%2-2)/2; i++ {
		fmt.Print("2 ")
	}
}
