package main

import (
	"fmt"
)

func watermelon() {
	var x int
	fmt.Scan(&x)
	if x%2 == 0 && x != 2 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}

}
