package main

import (
	"bufio"
	"fmt"
	"os"
)

func SquareString() {
	in := bufio.NewReader(os.Stdin)
	var t int
	var inp string
	var check bool
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		check = true
		fmt.Fscan(in, &inp)
		if len(inp)%2 != 0 {
			fmt.Println("NO")
		} else {
			for i := 0; i < len(inp)/2; i++ {
				if inp[i] != inp[i+len(inp)/2] {
					fmt.Println("NO")
					check = false
					break
				}
			}
			if check {
				fmt.Println("YES")
			}
		}

	}
}
