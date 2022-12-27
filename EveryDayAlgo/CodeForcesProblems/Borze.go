package main

import (
	"fmt"
)

func Borze() {
	var inp, res string
	fmt.Scan(&inp)
	for i := 0; i <= len(inp)-1; i++ {
		if inp[i] == '.' {
			res += "0"
		} else if inp[i+1] == '.' {
			res += "1"
			i += 1
		} else {
			res += "2"
			i += 1
		}
	}
	fmt.Print(res)
}
