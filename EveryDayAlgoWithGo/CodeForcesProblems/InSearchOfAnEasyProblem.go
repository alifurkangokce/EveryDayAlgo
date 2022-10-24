package main

import "fmt"

func InSearchOfAnEasyProblem() {
	var inp int
	fmt.Scan(&inp)
	for i := 0; i < inp; i++ {
		var inp2 int
		fmt.Scan(&inp2)
		if inp2 == 1 {
			fmt.Print("HARD")
			return
		}
	}
	fmt.Print("EASY")
}
