package main

import (
	"bufio"
	"fmt"
	"os"
)

func TwoElevators() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	var a, b, c int
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b, &c)
		if a == 1 {
			fmt.Println(1)
		} else {
			if b-c > 0 {
				if a-(b-c) == c {
					fmt.Println(3)
				} else if a-(b-c) > c {
					fmt.Println(2)
				} else {
					fmt.Println(1)
				}
			} else {
				if a+(b-c) == c {
					fmt.Println(3)
				} else if a+(b-c) > c {
					fmt.Println(2)
				} else {
					fmt.Println(1)
				}
			}
		}
	}
}
