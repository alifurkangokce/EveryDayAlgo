package main

import (
	"bufio"
	"fmt"
	"os"
)

func HQNinePlus() {
	in := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(in, &s)
	mm := make(map[rune]bool)
	mm['H'] = true
	mm['Q'] = true
	mm['9'] = true
	for _, v := range s {
		if mm[v] {
			fmt.Print("YES")
			return
		}
	}
	fmt.Print("NO")
}
