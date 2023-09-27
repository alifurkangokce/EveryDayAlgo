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
		if k == 1 && l == 1 {
			fmt.Println(1)
		} else {
			if k*l%2 != 0 {
				fmt.Println((k * l / 2) + 1)
			} else {
				fmt.Println((k * l / 2))
			}
		}

	}
}
