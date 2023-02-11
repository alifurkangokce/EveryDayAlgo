package main

import (
	"fmt"
)

var b2s = map[bool]string{false: "NO", true: "YES"}

func Lucky() {
	var i int
	var s string
	for fmt.Scanln(&i); i > 0; i-- {
		fmt.Scanln(&s)
		fmt.Println(b2s[s[0]+s[1]+s[2] == s[3]+s[4]+s[5]])
	}
}
