package main

import (
	"fmt"
)

func SoftDrinking() {
	var n, k, l, c, d, p, nl, np, a int
	fmt.Scan(&n, &k, &l, &c, &d, &p, &nl, &np)
	k *= l
	c *= d
	k /= nl
	p /= np
	if k > p {
		a = p
	} else {
		a = k
	}
	if a > c {
		a = c
	}
	fmt.Println(a / n)
}
