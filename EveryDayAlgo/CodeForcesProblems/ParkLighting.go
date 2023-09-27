package main

import (
	"bufio"
	"fmt"
	"os"
)

func ParkLighting() {
	in := bufio.NewReader(os.Stdin)
	var t, k, l int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &k, &l)
		fmt.Println(((k * l) + 1) / 2)

	}
}
