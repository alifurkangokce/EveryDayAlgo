package main

import "fmt"

func BoyOrGirl() {
	var inp string
	fmt.Scan(&inp)
	x := map[byte]int{}
	counter := 0
	for i := 0; i <= len(inp)-1; i++ {
		if x[inp[i]] == 0 {
			x[inp[i]]++
			counter++
		} else {
			x[inp[i]]++
		}
	}
	if counter%2 == 0 {
		fmt.Print("CHAT WITH HER!")
	} else {
		fmt.Print("IGNORE HIM!")
	}
}
