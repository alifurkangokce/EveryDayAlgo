package main

import (
	"bufio"
	"fmt"
	"os"
)

func Magnets() {
	var n int
	fmt.Scan(&n)
	var s, prev string
	buff := bufio.NewScanner(os.Stdin)
	res := 0
	for buff.Scan() {
		s = buff.Text()
		if s != prev {
			res += 1
		}
		prev = s
	}
	fmt.Println(res)
}
