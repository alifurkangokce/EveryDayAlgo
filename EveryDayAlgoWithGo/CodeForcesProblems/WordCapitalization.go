package main

import (
	"fmt"
	"strings"
)

func wordCapitalization() {
	var i string
	fmt.Scan(&i)
	s := strings.ToUpper(i[:1]) + i[1:]
	fmt.Print(s)
}
