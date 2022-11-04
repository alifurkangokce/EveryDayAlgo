package main

import (
	"fmt"
	"strings"
)

func Pangram() {
	var inp int
	var str string
	fmt.Scan(&inp, &str)
	str = strings.ToLower(str)
	if len(str) < 26 {
		fmt.Print("NO")
		return
	}
	m := make(map[byte]bool)
	for i := 0; i < inp; i++ {
		m[str[i]] = true
	}
	if len(m) == 26 {
		fmt.Printf("YES")
		return
	}
	fmt.Print("NO")

}
