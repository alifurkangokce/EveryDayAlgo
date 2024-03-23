package main

import (
	"bufio"
	"fmt"
	"os"
)

func VladandtheBestofFive() {
	in := bufio.NewReader(os.Stdin)
	var t int
	var s string
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &s)
		var aCount, bCount int
		for _, v := range s {
			if v == 'A' {
				aCount++
			} else {
				bCount++
			}
		}
		if aCount > bCount {
			fmt.Println("A")
		} else {
			fmt.Println("B")
		}
	}
}
