package main

import "fmt"

func AntonAndDanik() {
	var inp, d, a int
	var s string
	fmt.Scan(&inp, &s)
	for i := 0; i < inp; i++ {
		if s[i] == 'A' {
			a++
		} else {
			d++
		}
	}
	if d > a {
		fmt.Print("Danik")
		return
	} else if a > d {
		fmt.Print("Anton")
	} else {
		fmt.Print("Friendship")
	}

}
