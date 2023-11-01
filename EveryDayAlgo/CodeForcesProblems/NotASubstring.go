package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func NotASubString() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscanln(in, &t)
	r := strings.Repeat
	for ; t > 0; t-- {
		var w string
		fmt.Fscanln(in, &w)
		if w == "()" {
			fmt.Println("NO")
			continue
		}
		fmt.Println("YES")
		n := len(w)
		a := r("(", n) + r(")", n)
		b := r("()", n)
		if strings.Contains(b, w) {
			fmt.Println(a)
		} else {
			fmt.Println(b)
		}
	}
}
