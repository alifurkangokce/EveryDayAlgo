package main

import (
	"fmt"
	"strings"
)

func Word() {
	var inp string
	fmt.Scan(&inp)
	var u, l int
	for i := 0; i <= len(inp)-1; i++ {
		if inp[i] >= 65 && inp[i] <= 90 {
			u++
		} else {
			l++
		}
	}
	if l >= u {
		fmt.Print(strings.ToLower(inp))
	} else {
		fmt.Print(strings.ToUpper(inp))
	}
}
