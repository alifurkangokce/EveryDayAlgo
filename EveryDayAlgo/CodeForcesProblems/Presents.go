package main

import "fmt"

func Presents() {
	var friends int
	fmt.Scan(&friends)
	var reciever int
	gifts := make([]int, friends)
	for giver := 1; giver <= friends; giver++ {
		fmt.Scan(&reciever)
		gifts[reciever-1] = giver
	}

	for i := 0; i < friends; i++ {
		fmt.Print(gifts[i], " ")
	}
}
