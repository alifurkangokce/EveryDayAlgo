package main

import "fmt"

func NewYearAndHurry() {
	var inp1, inp2, cnt, i int
	fmt.Scan(&inp1, &inp2)
	res := 240 - inp2
	i = 1
	for res >= i*5 {
		cnt++
		res -= i * 5
		i++
	}
	if inp1 < cnt {
		fmt.Print(inp1)
	} else {
		fmt.Print(cnt)
	}

}
