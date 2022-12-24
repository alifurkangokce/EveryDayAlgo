package main

import "fmt"

func MishkaandGame() {
	var n, m, c, mp, cp int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m, &c)
		if m > c {
			mp++
		} else if c > m {
			cp++
		}
	}
	if mp > cp {
		fmt.Print("Mishka")
	} else if cp > mp {
		fmt.Print("Chris")
	} else {
		fmt.Print("Friendship is magic!^^")
	}
}
