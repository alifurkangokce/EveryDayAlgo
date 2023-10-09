package main

import (
	"bufio"
	"fmt"
	"os"
)

func ParkLighting() {
	in := bufio.NewReader(os.Stdin)
	var t, n, m int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &m)
		fmt.Println(((n * m) + 1) / 2)
	}
}
