package main

import (
	"bufio"
	"fmt"
	"os"
)

func TwoDivisors() {
	in := bufio.NewReader(os.Stdin)
	var T int
	fmt.Fscan(in, &T)
	for ; T > 0; T-- {
		var a, b int
		fmt.Fscan(in, &a, &b)
		solve(a, b)
	}
}

func solve(a int, b int) {
	res := 1
	if b < a {
		a, b = b, a
	}
	g := gcd(a, b)
	if g == a {
		res = b * b / g
	} else {
		res = b * a / g
	}
	fmt.Println(res)
	//fmt.Println()
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}
